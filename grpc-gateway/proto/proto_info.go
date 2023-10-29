package grpcgateway

import (
	"bytes"
	"fmt"
	"github.com/LCY2013/grpc-cloud/grpc-gateway/proto/internal"
	protov1 "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
	"math"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ReflectionProto knows how to format file descriptors as proto source code. Its fields
// provide some control over how the resulting source file is constructed and
// formatted.
type ReflectionProto struct {
	// If true, comments are rendered using "/*" style comments. Otherwise, they
	// are printed using "//" style line comments.
	PreferMultiLineStyleComments bool

	// If true, elements are sorted into a canonical order.
	//
	// The canonical order for elements in a file follows:
	//  1. Syntax
	//  2. Package
	//  3. Imports (sorted lexically)
	//  4. Options (sorted by name, standard options before custom options)
	//  5. Messages (sorted by name)
	//  6. Enums (sorted by name)
	//  7. Services (sorted by name)
	//  8. Extensions (grouped by extendee, sorted by extendee+tag)
	//
	// The canonical order of elements in a message follows:
	//  1. Options (sorted by name, standard options before custom options)
	//  2. Fields and One-Ofs (sorted by tag; one-ofs interleaved based on the
	//     minimum tag therein)
	//  3. Nested Messages (sorted by name)
	//  4. Nested Enums (sorted by name)
	//  5. Extension ranges (sorted by starting tag number)
	//  6. Nested Extensions (grouped by extendee, sorted by extendee+tag)
	//  7. Reserved ranges (sorted by starting tag number)
	//  8. Reserved names (sorted lexically)
	//
	// Methods are sorted within a service by name and appear after any service
	// options (which are sorted by name, standard options before custom ones).
	// Enum values are sorted within an enum, first by numeric value then by
	// name, and also appear after any enum options.
	//
	// Options for fields, enum values, and extension ranges are sorted by name,
	// standard options before custom ones.
	SortElements bool

	// The "less" function used to sort elements when printing. It is given two
	// elements, a and b, and should return true if a is "less than" b. In this
	// case, "less than" means that element a should appear earlier in the file
	// than element b.
	//
	// If this field is nil, no custom sorting is done and the SortElements
	// field is consulted to decide how to order the output. If this field is
	// non-nil, the SortElements field is ignored and this function is called to
	// order elements.
	CustomSortFunction func(a, b Element) bool

	// The indentation used. Any characters other than spaces or tabs will be
	// replaced with spaces. If unset/empty, two spaces will be used.
	Indent string

	// If true, detached comments (between elements) will be ignored.
	//
	// Deprecated: Use OmitComments bitmask instead.
	OmitDetachedComments bool

	// A bitmask of comment types to omit. If unset, all comments will be
	// included. Use CommentsAll to not print any comments.
	OmitComments CommentType

	// If true, trailing comments that typically appear on the same line as an
	// element (option, field, enum value, method) will be printed on a separate
	// line instead.
	//
	// So, with this set, you'll get output like so:
	//
	//    // leading comment for field
	//    repeated string names = 1;
	//    // trailing comment
	//
	// If left false, the printer will try to emit trailing comments on the same
	// line instead:
	//
	//    // leading comment for field
	//    repeated string names = 1; // trailing comment
	//
	// If the trailing comment has more than one line, it will automatically be
	// forced to the next line.
	TrailingCommentsOnSeparateLine bool

	// If true, the printed output will eschew any blank lines, which otherwise
	// appear between descriptor elements and comment blocks. Note that if
	// detached comments are being printed, this will cause them to be merged
	// into the subsequent leading comments. Similarly, any element trailing
	// comments will be merged into the subsequent leading comments.
	Compact bool

	// If true, all references to messages, extensions, and enums (such as in
	// options, field types, and method request and response types) will be
	// fully-qualified. When left unset, the referenced elements will contain
	// only as much qualifier as is required.
	//
	// For example, if a message is in the same package as the reference, the
	// simple name can be used. If a message shares some context with the
	// reference, only the unshared context needs to be included. For example:
	//
	//  message Foo {
	//    message Bar {
	//      enum Baz {
	//        ZERO = 0;
	//        ONE = 1;
	//      }
	//    }
	//
	//    // This field shares some context as the enum it references: they are
	//    // both inside of the namespace Foo:
	//    //    field is "Foo.my_baz"
	//    //     enum is "Foo.Bar.Baz"
	//    // So we only need to qualify the reference with the context that they
	//    // do NOT have in common:
	//    Bar.Baz my_baz = 1;
	//  }
	//
	// When printing fully-qualified names, they will be preceded by a dot, to
	// avoid any ambiguity that they might be relative vs. fully-qualified.
	ForceFullyQualifiedNames bool

	// The number of options that trigger short options expressions to be
	// rendered using multiple lines. Short options expressions are those
	// found on fields and enum values, that use brackets ("[" and "]") and
	// comma-separated options. If more options than this are present, they
	// will be expanded to multiple lines (one option per line).
	//
	// If unset (e.g. if zero), a default threshold of 3 is used.
	ShortOptionsExpansionThresholdCount int

	// The length of printed options that trigger short options expressions to
	// be rendered using multiple lines. If the short options contain more than
	// one option and their printed length is longer than this threshold, they
	// will be expanded to multiple lines (one option per line).
	//
	// If unset (e.g. if zero), a default threshold of 50 is used.
	ShortOptionsExpansionThresholdLength int

	// The length of a printed option value message literal that triggers the
	// message literal to be rendered using multiple lines instead of using a
	// compact single-line form. The message must include at least two fields
	// or contain a field that is a nested message to be expanded.
	//
	// This value is further used to decide when to expand individual field
	// values that are nested message literals or array literals (for repeated
	// fields).
	//
	// If unset (e.g. if zero), a default threshold of 50 is used.
	MessageLiteralExpansionThresholdLength int

	// Syntax proto version
	Syntax string

	// Package proto package
	Package string

	// Dependency dep
	Dependency []string

	// Comment
	Comment []string

	// FileOptions
	FileOptions []option

	// FileMessageDescriptor
	FileMessageDescriptor []*desc.MessageDescriptor

	// FileEnumDescriptor
	FileEnumDescriptor []*desc.EnumDescriptor

	// FileServiceInfo
	FileServiceInfo []*ServiceInfo

	// FileFieldDescriptor
	FileFieldDescriptor []*desc.FieldDescriptor

	// FileExtensions
	FileExtensions extensions

	// MethodOptions
	MethodOptions map[*desc.MethodDescriptor][]option
}

// ServiceInfo rpc info
type ServiceInfo struct {
	ServiceDescriptor *desc.ServiceDescriptor
	MethodDescriptor  []*desc.MethodDescriptor
	Options           []option
}

// CommentType is a kind of comments in a proto source file. This can be used
// as a bitmask.
type CommentType int

const (
	// CommentsDetached refers to comments that are not "attached" to any
	// source element. They are attributed to the subsequent element in the
	// file as "detached" comments.
	CommentsDetached CommentType = 1 << iota
	// CommentsTrailing refers to a comment block immediately following an
	// element in the source file. If another element immediately follows
	// the trailing comment, it is instead considered a leading comment for
	// that subsequent element.
	CommentsTrailing
	// CommentsLeading refers to a comment block immediately preceding an
	// element in the source file. For high-level elements (those that have
	// their own descriptor), these are used as doc comments for that element.
	CommentsLeading
	// CommentsTokens refers to any comments (leading, trailing, or detached)
	// on low-level elements in the file. "High-level" elements have their own
	// descriptors, e.g. messages, enums, fields, services, and methods. But
	// comments can appear anywhere (such as around identifiers and keywords,
	// sprinkled inside the declarations of a high-level element). This class
	// of comments are for those extra comments sprinkled into the file.
	CommentsTokens

	// CommentsNonDoc refers to comments that are *not* doc comments. This is a
	// bitwise union of everything other than CommentsLeading. If you configure
	// a printer to omit this, only doc comments on descriptor elements will be
	// included in the printed output.
	CommentsNonDoc = CommentsDetached | CommentsTrailing | CommentsTokens
	// CommentsAll indicates all kinds of comments. If you configure a printer
	// to omit this, no comments will appear in the printed output, even if the
	// input descriptors had source info and comments.
	CommentsAll = -1
)

// pkg represents a package name
type pkg string

// imp represents an imported file name
type imp string

// ident represents an identifier
type ident string

// messageVal represents a message value for an option
type messageVal struct {
	// the package and scope in which the option value is defined
	pkg, scope string
	// the option value
	msg proto.Message
}

// option represents a resolved descriptor option
type option struct {
	name string
	val  interface{}
}

// reservedRange represents a reserved range from a message or enum
type reservedRange struct {
	start, end int32
}

type edgeKind int

const (
	edgeKindOption edgeKind = iota
	edgeKindFile
	edgeKindMessage
	edgeKindField
	edgeKindOneOf
	edgeKindExtensionRange
	edgeKindReservedRange
	edgeKindReservedName
	edgeKindEnum
	edgeKindEnumVal
	edgeKindService
	edgeKindMethod
)

// edges in simple state machine for matching options paths
// whose prefix should be included in source info to handle
// the way options are printed (which cannot always include
// the full path from original source)
var edges = map[edgeKind]map[int32]edgeKind{
	edgeKindFile: {
		internal.File_optionsTag:    edgeKindOption,
		internal.File_messagesTag:   edgeKindMessage,
		internal.File_enumsTag:      edgeKindEnum,
		internal.File_extensionsTag: edgeKindField,
		internal.File_servicesTag:   edgeKindService,
	},
	edgeKindMessage: {
		internal.Message_optionsTag:        edgeKindOption,
		internal.Message_fieldsTag:         edgeKindField,
		internal.Message_oneOfsTag:         edgeKindOneOf,
		internal.Message_nestedMessagesTag: edgeKindMessage,
		internal.Message_enumsTag:          edgeKindEnum,
		internal.Message_extensionsTag:     edgeKindField,
		internal.Message_extensionRangeTag: edgeKindExtensionRange,
		internal.Message_reservedRangeTag:  edgeKindReservedRange,
		internal.Message_reservedNameTag:   edgeKindReservedName,
	},
	edgeKindField: {
		internal.Field_optionsTag: edgeKindOption,
	},
	edgeKindOneOf: {
		internal.OneOf_optionsTag: edgeKindOption,
	},
	edgeKindExtensionRange: {
		internal.ExtensionRange_optionsTag: edgeKindOption,
	},
	edgeKindEnum: {
		internal.Enum_optionsTag:       edgeKindOption,
		internal.Enum_valuesTag:        edgeKindEnumVal,
		internal.Enum_reservedRangeTag: edgeKindReservedRange,
		internal.Enum_reservedNameTag:  edgeKindReservedName,
	},
	edgeKindEnumVal: {
		internal.EnumVal_optionsTag: edgeKindOption,
	},
	edgeKindService: {
		internal.Service_optionsTag: edgeKindOption,
		internal.Service_methodsTag: edgeKindMethod,
	},
	edgeKindMethod: {
		internal.Method_optionsTag: edgeKindOption,
	},
}

// ProtoInfoDesc desc proto info
func ProtoInfoDesc(dsc desc.Descriptor) *ReflectionProto {
	rp := &ReflectionProto{
		Compact:                  true,
		OmitComments:             CommentsNonDoc,
		SortElements:             true,
		ForceFullyQualifiedNames: true,
		MethodOptions:            make(map[*desc.MethodDescriptor][]option),
	}

	/*fdp := dsc.GetFile().AsFileDescriptorProto()
	sourceInfo := internal.CreateSourceInfoMap(fdp)
	extendOptionLocations(sourceInfo, fdp.GetSourceCodeInfo().GetLocation())

	var reg protoregistry.Types
	internal.RegisterTypesForFile(&reg, dsc.GetFile().UnwrapFile())
	reparseUnknown(&reg, fdp.ProtoReflect())*/
	er := dynamic.ExtensionRegistry{}
	er.AddExtensionsFromFileRecursively(dsc.GetFile())
	reg := dynamic.NewMessageFactoryWithExtensionRegistry(&er)
	fdp := dsc.GetFile().AsFileDescriptorProto()
	sourceInfo := internal.CreateSourceInfoMap(fdp)
	extendOptionLocations(sourceInfo, fdp.GetSourceCodeInfo().GetLocation())

	path := findElement(dsc)
	switch d := dsc.(type) {
	case *desc.FileDescriptor:
		rp.descFile(d, reg, sourceInfo)
	case *desc.MessageDescriptor:
		rp.descMessage(d, reg, sourceInfo, path, 0)
	case *desc.FieldDescriptor:
		var scope string
		if md, ok := d.GetParent().(*desc.MessageDescriptor); ok {
			scope = md.GetFullyQualifiedName()
		} else {
			scope = d.GetFile().GetPackage()
		}
		if d.IsExtension() {
			extNameSi := sourceInfo.Get(append(path, internal.Field_extendeeTag))
			rp.descElementString(extNameSi, 0, rp.qualifyName(d.GetFile().GetPackage(), scope, d.GetOwner().GetFullyQualifiedName()))

			rp.descField(d, reg, sourceInfo, path, scope, 1)

		} else {
			rp.descField(d, reg, sourceInfo, path, scope, 0)
		}
	case *desc.OneOfDescriptor:
		md := d.GetOwner()
		elements := elementAddrs{dsc: md}
		for i := range md.GetFields() {
			elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_fieldsTag, elementIndex: i})
		}
		rp.descOneOf(d, elements, 0, reg, sourceInfo, path[:len(path)-1], 0, path[len(path)-1])
	case *desc.EnumDescriptor:
		rp.descEnum(d, reg, sourceInfo, path, 0)
	case *desc.EnumValueDescriptor:
		rp.descEnumValue(d, reg, sourceInfo, path, 0)
	case *desc.ServiceDescriptor:
		rp.descService(d, reg, sourceInfo, path, 0)
	case *desc.MethodDescriptor:
		rp.descMethod(d, reg, sourceInfo, path, 0)
	}

	return rp
}

func extendOptionLocations(sc internal.SourceInfoMap, locs []*descriptorpb.SourceCodeInfo_Location) {
	// we iterate in the order that locations appear in descriptor
	// for determinism (if we ranged over the map, order and thus
	// potentially results are non-deterministic)
	for _, loc := range locs {
		allowed := edges[edgeKindFile]
		for i := 0; i+1 < len(loc.Path); i += 2 {
			nextKind, ok := allowed[loc.Path[i]]
			if !ok {
				break
			}
			if nextKind == edgeKindOption {
				// We've found an option entry. This could be arbitrarily deep
				// (for options that are nested messages) or it could end
				// abruptly (for non-repeated fields). But we need a path that
				// is exactly the path-so-far plus two: the option tag and an
				// optional index for repeated option fields (zero for
				// non-repeated option fields). This is used for querying source
				// info when printing options.
				newPath := make([]int32, i+3)
				copy(newPath, loc.Path)
				sc.PutIfAbsent(newPath, loc)
				// we do another path of path-so-far plus two, but with
				// explicit zero index -- just in case this actual path has
				// an extra path element, but it's not an index (e.g the
				// option field is not repeated, but the source info we are
				// looking at indicates a tag of a nested field)
				newPath[len(newPath)-1] = 0
				sc.PutIfAbsent(newPath, loc)
				// finally, we need the path-so-far plus one, just the option
				// tag, for sorting option groups
				newPath = newPath[:len(newPath)-1]
				sc.PutIfAbsent(newPath, loc)

				break
			} else {
				allowed = edges[nextKind]
			}
		}
	}
}

func reparseUnknown(reg *protoregistry.Types, msg protoreflect.Message) {
	msg.Range(func(fld protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		if fld.Kind() != protoreflect.MessageKind && fld.Kind() != protoreflect.GroupKind {
			return true
		}
		if fld.IsList() {
			l := val.List()
			for i := 0; i < l.Len(); i++ {
				reparseUnknown(reg, l.Get(i).Message())
			}
		} else if fld.IsMap() {
			mapVal := fld.MapValue()
			if mapVal.Kind() != protoreflect.MessageKind && mapVal.Kind() != protoreflect.GroupKind {
				return true
			}
			m := val.Map()
			m.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				reparseUnknown(reg, v.Message())
				return true
			})
		} else {
			reparseUnknown(reg, val.Message())
		}
		return true
	})

	unk := msg.GetUnknown()
	if len(unk) > 0 {
		other := msg.New().Interface()
		if err := (proto.UnmarshalOptions{Resolver: reg}).Unmarshal(unk, other); err == nil {
			msg.SetUnknown(nil)
			proto.Merge(msg.Interface(), other)
		}
	}
}

func findElement(dsc desc.Descriptor) []int32 {
	if dsc.GetParent() == nil {
		return nil
	}
	path := findElement(dsc.GetParent())
	switch d := dsc.(type) {
	case *desc.MessageDescriptor:
		if pm, ok := d.GetParent().(*desc.MessageDescriptor); ok {
			return append(path, internal.Message_nestedMessagesTag, getMessageIndex(d, pm.GetNestedMessageTypes()))
		}
		return append(path, internal.File_messagesTag, getMessageIndex(d, d.GetFile().GetMessageTypes()))

	case *desc.FieldDescriptor:
		if d.IsExtension() {
			if pm, ok := d.GetParent().(*desc.MessageDescriptor); ok {
				return append(path, internal.Message_extensionsTag, getFieldIndex(d, pm.GetNestedExtensions()))
			}
			return append(path, internal.File_extensionsTag, getFieldIndex(d, d.GetFile().GetExtensions()))
		}
		return append(path, internal.Message_fieldsTag, getFieldIndex(d, d.GetOwner().GetFields()))

	case *desc.OneOfDescriptor:
		return append(path, internal.Message_oneOfsTag, getOneOfIndex(d, d.GetOwner().GetOneOfs()))

	case *desc.EnumDescriptor:
		if pm, ok := d.GetParent().(*desc.MessageDescriptor); ok {
			return append(path, internal.Message_enumsTag, getEnumIndex(d, pm.GetNestedEnumTypes()))
		}
		return append(path, internal.File_enumsTag, getEnumIndex(d, d.GetFile().GetEnumTypes()))

	case *desc.EnumValueDescriptor:
		return append(path, internal.Enum_valuesTag, getEnumValueIndex(d, d.GetEnum().GetValues()))

	case *desc.ServiceDescriptor:
		return append(path, internal.File_servicesTag, getServiceIndex(d, d.GetFile().GetServices()))

	case *desc.MethodDescriptor:
		return append(path, internal.Service_methodsTag, getMethodIndex(d, d.GetService().GetMethods()))

	default:
		panic(fmt.Sprintf("unexpected descriptor type: %T", dsc))
	}
}

func getMessageIndex(md *desc.MessageDescriptor, list []*desc.MessageDescriptor) int32 {
	for i := range list {
		if md == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of message %s", md.GetFullyQualifiedName()))
}

func getFieldIndex(fd *desc.FieldDescriptor, list []*desc.FieldDescriptor) int32 {
	for i := range list {
		if fd == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of field %s", fd.GetFullyQualifiedName()))
}

func getOneOfIndex(ood *desc.OneOfDescriptor, list []*desc.OneOfDescriptor) int32 {
	for i := range list {
		if ood == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of oneof %s", ood.GetFullyQualifiedName()))
}

func getEnumIndex(ed *desc.EnumDescriptor, list []*desc.EnumDescriptor) int32 {
	for i := range list {
		if ed == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of enum %s", ed.GetFullyQualifiedName()))
}

func getEnumValueIndex(evd *desc.EnumValueDescriptor, list []*desc.EnumValueDescriptor) int32 {
	for i := range list {
		if evd == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of enum value %s", evd.GetFullyQualifiedName()))
}

func getServiceIndex(sd *desc.ServiceDescriptor, list []*desc.ServiceDescriptor) int32 {
	for i := range list {
		if sd == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of service %s", sd.GetFullyQualifiedName()))
}

func getMethodIndex(mtd *desc.MethodDescriptor, list []*desc.MethodDescriptor) int32 {
	for i := range list {
		if mtd == list[i] {
			return int32(i)
		}
	}
	panic(fmt.Sprintf("unable to determine index of method %s", mtd.GetFullyQualifiedName()))
}

func findExtSi(fieldSi *descriptorpb.SourceCodeInfo_Location, extSis []*descriptorpb.SourceCodeInfo_Location) *descriptorpb.SourceCodeInfo_Location {
	if len(fieldSi.GetSpan()) == 0 {
		return nil
	}
	for _, extSi := range extSis {
		if isSpanWithin(fieldSi.Span, extSi.Span) {
			return extSi
		}
	}
	return nil
}

func isSpanWithin(span, enclosing []int32) bool {
	start := enclosing[0]
	var end int32
	if len(enclosing) == 3 {
		end = enclosing[0]
	} else {
		end = enclosing[2]
	}
	if span[0] < start || span[0] > end {
		return false
	}

	if span[0] == start {
		return span[1] >= enclosing[1]
	} else if span[0] == end {
		return span[1] <= enclosing[len(enclosing)-1]
	}
	return true
}

type extensionDecl struct {
	extendee   string
	sourceInfo *descriptorpb.SourceCodeInfo_Location
	fields     []*desc.FieldDescriptor
}

type extensions map[*desc.FieldDescriptor]*extensionDecl

func valueForOption(fld protoreflect.FieldDescriptor, val interface{}) interface{} {
	switch val := val.(type) {
	case protoreflect.EnumNumber:
		ev := fld.Enum().Values().ByNumber(val)
		if ev == nil {
			// if enum val is unknown, we'll return nil and have to skip it :(
			return nil
		}
		return ident(ev.Name())
	case protoreflect.Message:
		return val.Interface()
	default:
		return val
	}
}

func toUninterpretedOption(message proto.Message) *descriptorpb.UninterpretedOption {
	if uo, ok := message.(*descriptorpb.UninterpretedOption); ok {
		return uo
	}
	// marshal and unmarshal to convert; if we fail to convert, skip it
	var uo descriptorpb.UninterpretedOption
	data, err := proto.Marshal(message)
	if err != nil {
		return nil
	}
	if proto.Unmarshal(data, &uo) != nil {
		return nil
	}
	return &uo
}

func uninterpretedToOptions(uninterp []*descriptorpb.UninterpretedOption) []option {
	opts := make([]option, len(uninterp))
	for i, unint := range uninterp {
		var buf bytes.Buffer
		for ni, n := range unint.Name {
			if ni > 0 {
				buf.WriteByte('.')
			}
			if n.GetIsExtension() {
				_, _ = fmt.Fprintf(&buf, "(%s)", n.GetNamePart())
			} else {
				buf.WriteString(n.GetNamePart())
			}
		}

		var v interface{}
		switch {
		case unint.IdentifierValue != nil:
			v = ident(unint.GetIdentifierValue())
		case unint.StringValue != nil:
			v = string(unint.GetStringValue())
		case unint.DoubleValue != nil:
			v = unint.GetDoubleValue()
		case unint.PositiveIntValue != nil:
			v = unint.GetPositiveIntValue()
		case unint.NegativeIntValue != nil:
			v = unint.GetNegativeIntValue()
		case unint.AggregateValue != nil:
			v = ident(unint.GetAggregateValue())
		}

		opts[i] = option{name: buf.String(), val: v}
	}
	return opts
}

func optionsAsElementAddrs(optionsTag int32, order int, opts map[int32][]option) []elementAddr {
	optAddrs := make([]elementAddr, 0, len(opts))
	for tag := range opts {
		optAddrs = append(optAddrs, elementAddr{elementType: optionsTag, elementIndex: int(tag), order: order})
	}
	// We want stable output. So, if the printer can't sort these a better way,
	// they'll at least be in a deterministic order (by name).
	sort.Sort(optionsByName{addrs: optAddrs, opts: opts})
	return optAddrs
}

// quotedBytes implements the text format for string literals for protocol
// buffers. Since the underlying data is a bytes field, this encodes all
// bytes outside the 7-bit ASCII printable range. To preserve unicode strings
// without byte escapes, use quotedString.
func quotedBytes(s string) string {
	var b bytes.Buffer
	b.WriteByte('"')
	// Loop over the bytes, not the runes.
	for i := 0; i < len(s); i++ {
		// Divergence from C++: we don't escape apostrophes.
		// There's no need to escape them, and the C++ parser
		// copes with a naked apostrophe.
		switch c := s[i]; c {
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		case '"':
			b.WriteString("\\")
		case '\\':
			b.WriteString("\\\\")
		default:
			if c >= 0x20 && c < 0x7f {
				b.WriteByte(c)
			} else {
				_, _ = fmt.Fprintf(&b, "\\%03o", c)
			}
		}
	}
	b.WriteByte('"')

	return b.String()
}

// quotedString implements the text format for string literals for protocol
// buffers. This form is also acceptable for string literals in option values
// by the protocol buffer compiler, protoc.
func quotedString(s string) string {
	var b bytes.Buffer
	b.WriteByte('"')
	// Loop over the bytes, not the runes.
	for {
		r, n := utf8.DecodeRuneInString(s)
		if n == 0 {
			break // end of string
		}
		if r == utf8.RuneError && n == 1 {
			// Invalid UTF8! Use an octal byte escape to encode the bad byte.
			_, _ = fmt.Fprintf(&b, "\\%03o", s[0])
			s = s[1:]
			continue
		}

		// Divergence from C++: we don't escape apostrophes.
		// There's no need to escape them, and the C++ parser
		// copes with a naked apostrophe.
		switch r {
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		case '"':
			b.WriteString("\\")
		case '\\':
			b.WriteString("\\\\")
		default:
			if unicode.IsPrint(r) {
				b.WriteRune(r)
			} else {
				// if it's not printable, use a unicode escape
				if r > 0xffff {
					_, _ = fmt.Fprintf(&b, "\\U%08X", r)
				} else if r > 0x7F {
					_, _ = fmt.Fprintf(&b, "\\u%04X", r)
				} else {
					_, _ = fmt.Fprintf(&b, "\\%03o", byte(r))
				}
			}
		}

		s = s[n:]
	}

	b.WriteByte('"')

	return b.String()
}

type elementAddr struct {
	elementType  int32
	elementIndex int
	order        int
}

type elementAddrs struct {
	addrs []elementAddr
	dsc   interface{}
	opts  map[int32][]option
}

func (a elementAddrs) Len() int {
	return len(a.addrs)
}

func (a elementAddrs) Less(i, j int) bool {
	// explicit order is considered first
	if a.addrs[i].order < a.addrs[j].order {
		return true
	} else if a.addrs[i].order > a.addrs[j].order {
		return false
	}
	// if order is equal, sort by element type
	if a.addrs[i].elementType < a.addrs[j].elementType {
		return true
	} else if a.addrs[i].elementType > a.addrs[j].elementType {
		return false
	}

	di := a.at(a.addrs[i])
	dj := a.at(a.addrs[j])

	switch vi := di.(type) {
	case *desc.FieldDescriptor:
		// fields are ordered by tag number
		vj := dj.(*desc.FieldDescriptor)
		// regular fields before extensions; extensions grouped by extendee
		if !vi.IsExtension() && vj.IsExtension() {
			return true
		} else if vi.IsExtension() && !vj.IsExtension() {
			return false
		} else if vi.IsExtension() && vj.IsExtension() {
			if vi.GetOwner() != vj.GetOwner() {
				return vi.GetOwner().GetFullyQualifiedName() < vj.GetOwner().GetFullyQualifiedName()
			}
		}
		return vi.GetNumber() < vj.GetNumber()

	case *desc.EnumValueDescriptor:
		// enum values ordered by number then name
		vj := dj.(*desc.EnumValueDescriptor)
		if vi.GetNumber() == vj.GetNumber() {
			return vi.GetName() < vj.GetName()
		}
		return vi.GetNumber() < vj.GetNumber()

	case *descriptorpb.DescriptorProto_ExtensionRange:
		// extension ranges ordered by tag
		return vi.GetStart() < dj.(*descriptorpb.DescriptorProto_ExtensionRange).GetStart()

	case reservedRange:
		// reserved ranges ordered by tag, too
		return vi.start < dj.(reservedRange).start

	case string:
		// reserved names lexically sorted
		return vi < dj.(string)

	case pkg:
		// reserved names lexically sorted
		return vi < dj.(pkg)

	case imp:
		// reserved names lexically sorted
		return vi < dj.(imp)

	case []option:
		// options sorted by name, extensions last
		return optionLess(vi, dj.([]option))

	default:
		// all other descriptors ordered by name
		return di.(desc.Descriptor).GetName() < dj.(desc.Descriptor).GetName()
	}
}

func (a elementAddrs) Swap(i, j int) {
	a.addrs[i], a.addrs[j] = a.addrs[j], a.addrs[i]
}

func (a elementAddrs) at(addr elementAddr) interface{} {
	switch dsc := a.dsc.(type) {
	case *desc.FileDescriptor:
		switch addr.elementType {
		case internal.File_packageTag:
			return pkg(dsc.GetPackage())
		case internal.File_dependencyTag:
			return imp(dsc.AsFileDescriptorProto().GetDependency()[addr.elementIndex])
		case internal.File_optionsTag:
			return a.opts[int32(addr.elementIndex)]
		case internal.File_messagesTag:
			return dsc.GetMessageTypes()[addr.elementIndex]
		case internal.File_enumsTag:
			return dsc.GetEnumTypes()[addr.elementIndex]
		case internal.File_servicesTag:
			return dsc.GetServices()[addr.elementIndex]
		case internal.File_extensionsTag:
			return dsc.GetExtensions()[addr.elementIndex]
		}
	case *desc.MessageDescriptor:
		switch addr.elementType {
		case internal.Message_optionsTag:
			return a.opts[int32(addr.elementIndex)]
		case internal.Message_fieldsTag:
			return dsc.GetFields()[addr.elementIndex]
		case internal.Message_nestedMessagesTag:
			return dsc.GetNestedMessageTypes()[addr.elementIndex]
		case internal.Message_enumsTag:
			return dsc.GetNestedEnumTypes()[addr.elementIndex]
		case internal.Message_extensionsTag:
			return dsc.GetNestedExtensions()[addr.elementIndex]
		case internal.Message_extensionRangeTag:
			return dsc.AsDescriptorProto().GetExtensionRange()[addr.elementIndex]
		case internal.Message_reservedRangeTag:
			rng := dsc.AsDescriptorProto().GetReservedRange()[addr.elementIndex]
			return reservedRange{start: rng.GetStart(), end: rng.GetEnd() - 1}
		case internal.Message_reservedNameTag:
			return dsc.AsDescriptorProto().GetReservedName()[addr.elementIndex]
		}
	case *desc.FieldDescriptor:
		if addr.elementType == internal.Field_optionsTag {
			return a.opts[int32(addr.elementIndex)]
		}
	case *desc.OneOfDescriptor:
		switch addr.elementType {
		case internal.OneOf_optionsTag:
			return a.opts[int32(addr.elementIndex)]
		case -internal.Message_fieldsTag:
			return dsc.GetOwner().GetFields()[addr.elementIndex]
		}
	case *desc.EnumDescriptor:
		switch addr.elementType {
		case internal.Enum_optionsTag:
			return a.opts[int32(addr.elementIndex)]
		case internal.Enum_valuesTag:
			return dsc.GetValues()[addr.elementIndex]
		case internal.Enum_reservedRangeTag:
			rng := dsc.AsEnumDescriptorProto().GetReservedRange()[addr.elementIndex]
			return reservedRange{start: rng.GetStart(), end: rng.GetEnd()}
		case internal.Enum_reservedNameTag:
			return dsc.AsEnumDescriptorProto().GetReservedName()[addr.elementIndex]
		}
	case *desc.EnumValueDescriptor:
		if addr.elementType == internal.EnumVal_optionsTag {
			return a.opts[int32(addr.elementIndex)]
		}
	case *desc.ServiceDescriptor:
		switch addr.elementType {
		case internal.Service_optionsTag:
			return a.opts[int32(addr.elementIndex)]
		case internal.Service_methodsTag:
			return dsc.GetMethods()[addr.elementIndex]
		}
	case *desc.MethodDescriptor:
		if addr.elementType == internal.Method_optionsTag {
			return a.opts[int32(addr.elementIndex)]
		}
	case extensionRange:
		if addr.elementType == internal.ExtensionRange_optionsTag {
			return a.opts[int32(addr.elementIndex)]
		}
	}

	panic(fmt.Sprintf("location for unknown field %d of %T", addr.elementType, a.dsc))
}

type extensionRange struct {
	owner    *desc.MessageDescriptor
	extRange *descriptorpb.DescriptorProto_ExtensionRange
}

type elementSrcOrder struct {
	elementAddrs
	sourceInfo internal.SourceInfoMap
	prefix     []int32
}

func (a elementSrcOrder) Less(i, j int) bool {
	ti := a.addrs[i].elementType
	ei := a.addrs[i].elementIndex

	tj := a.addrs[j].elementType
	ej := a.addrs[j].elementIndex

	var si, sj *descriptorpb.SourceCodeInfo_Location
	if ei < 0 {
		si = a.sourceInfo.Get(append(a.prefix, -int32(ei)))
	} else if ti < 0 {
		p := make([]int32, len(a.prefix)-2)
		copy(p, a.prefix)
		si = a.sourceInfo.Get(append(p, ti, int32(ei)))
	} else {
		si = a.sourceInfo.Get(append(a.prefix, ti, int32(ei)))
	}
	if ej < 0 {
		sj = a.sourceInfo.Get(append(a.prefix, -int32(ej)))
	} else if tj < 0 {
		p := make([]int32, len(a.prefix)-2)
		copy(p, a.prefix)
		sj = a.sourceInfo.Get(append(p, tj, int32(ej)))
	} else {
		sj = a.sourceInfo.Get(append(a.prefix, tj, int32(ej)))
	}

	if (si == nil) != (sj == nil) {
		// generally, we put unknown elements after known ones;
		// except package, imports, and option elements go first

		// i will be unknown and j will be known
		swapped := false
		if si != nil {
			ti, tj = tj, ti
			swapped = true
		}
		switch a.dsc.(type) {
		case *desc.FileDescriptor:
			// NB: These comparisons are *trying* to get things ordered so that
			// 1) If the package element has no source info, it appears _first_.
			// 2) If any import element has no source info, it appears _after_
			//    the package element but _before_ any other element.
			// 3) If any option element has no source info, it appears _after_
			//    the package and import elements but _before_ any other element.
			// If the package, imports, and options are all missing source info,
			// this will sort them all to the top in expected order. But if they
			// are mixed (some _do_ have source info, some do not), and elements
			// with source info have spans that positions them _after_ other
			// elements in the file, then this Less function will be unstable
			// since the above dual objectives for imports and options ("before
			// this but after that") may be in conflict with one another. This
			// should not cause any problems, other than elements being possibly
			// sorted in a confusing order.
			//
			// Well-formed descriptors should instead have consistent source
			// info: either all elements have source info or none do. So this
			// should not be an issue in practice.
			if ti == internal.File_packageTag {
				return !swapped
			}
			if ti == internal.File_dependencyTag {
				if tj == internal.File_packageTag {
					// imports will come *after* package
					return swapped
				}
				return !swapped
			}
			if ti == internal.File_optionsTag {
				if tj == internal.File_packageTag || tj == internal.File_dependencyTag {
					// options will come *after* package and imports
					return swapped
				}
				return !swapped
			}
		case *desc.MessageDescriptor:
			if ti == internal.Message_optionsTag {
				return !swapped
			}
		case *desc.EnumDescriptor:
			if ti == internal.Enum_optionsTag {
				return !swapped
			}
		case *desc.ServiceDescriptor:
			if ti == internal.Service_optionsTag {
				return !swapped
			}
		}
		return swapped

	} else if si == nil || sj == nil {
		// let stable sort keep unknown elements in same relative order
		return false
	}

	for idx := 0; idx < len(sj.Span); idx++ {
		if idx >= len(si.Span) {
			return true
		}
		if si.Span[idx] < sj.Span[idx] {
			return true
		}
		if si.Span[idx] > sj.Span[idx] {
			return false
		}
	}
	return false
}

type customSortOrder struct {
	elementAddrs
	less func(a, b Element) bool
}

func (cso customSortOrder) Less(i, j int) bool {
	ei := asElement(cso.at(cso.addrs[i]))
	ej := asElement(cso.at(cso.addrs[j]))
	return cso.less(ei, ej)
}

type optionsByName struct {
	addrs []elementAddr
	opts  map[int32][]option
}

func (o optionsByName) Len() int {
	return len(o.addrs)
}

func (o optionsByName) Less(i, j int) bool {
	oi := o.opts[int32(o.addrs[i].elementIndex)]
	oj := o.opts[int32(o.addrs[j].elementIndex)]
	return optionLess(oi, oj)
}

func (o optionsByName) Swap(i, j int) {
	o.addrs[i], o.addrs[j] = o.addrs[j], o.addrs[i]
}

func optionLess(i, j []option) bool {
	ni := i[0].name
	nj := j[0].name
	if ni[0] != '(' && nj[0] == '(' {
		return true
	} else if ni[0] == '(' && nj[0] != '(' {
		return false
	}
	return ni < nj
}

func (p *ReflectionProto) computeExtensions(sourceInfo internal.SourceInfoMap, exts []*desc.FieldDescriptor, path []int32) extensions {
	extsMap := map[string]map[*descriptorpb.SourceCodeInfo_Location]*extensionDecl{}
	extSis := sourceInfo.GetAll(path)
	for _, extd := range exts {
		name := extd.GetOwner().GetFullyQualifiedName()
		extSi := findExtSi(extd.GetSourceInfo(), extSis)
		extsBySi := extsMap[name]
		if extsBySi == nil {
			extsBySi = map[*descriptorpb.SourceCodeInfo_Location]*extensionDecl{}
			extsMap[name] = extsBySi
		}
		extDecl := extsBySi[extSi]
		if extDecl == nil {
			extDecl = &extensionDecl{
				sourceInfo: extSi,
				extendee:   name,
			}
			extsBySi[extSi] = extDecl
		}
		extDecl.fields = append(extDecl.fields, extd)
	}

	ret := extensions{}
	for _, extsBySi := range extsMap {
		for _, extDecl := range extsBySi {
			for _, extd := range extDecl.fields {
				ret[extd] = extDecl
			}
		}
	}
	return ret
}

func sortKeys(m map[interface{}]interface{}) []interface{} {
	res := make(sortedKeys, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}
	sort.Sort(res)
	return ([]interface{})(res)
}

type sortedKeys []interface{}

func (k sortedKeys) Len() int {
	return len(k)
}

func (k sortedKeys) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k sortedKeys) Less(i, j int) bool {
	switch i := k[i].(type) {
	case int32:
		return i < k[j].(int32)
	case uint32:
		return i < k[j].(uint32)
	case int64:
		return i < k[j].(int64)
	case uint64:
		return i < k[j].(uint64)
	case string:
		return i < k[j].(string)
	case bool:
		return !i && k[j].(bool)
	default:
		panic(fmt.Sprintf("invalid type for map key: %T", i))
	}
}

func (p *ReflectionProto) sort(elements elementAddrs, sourceInfo internal.SourceInfoMap, path []int32) {
	if p.CustomSortFunction != nil {
		sort.Stable(customSortOrder{elementAddrs: elements, less: p.CustomSortFunction})
	} else if p.SortElements {
		// canonical sorted order
		sort.Stable(elements)
	} else {
		// use source order (per location information in SourceCodeInfo); or
		// if that isn't present use declaration order, but grouped by type
		sort.Stable(elementSrcOrder{
			elementAddrs: elements,
			sourceInfo:   sourceInfo,
			prefix:       path,
		})
	}
}

func (p *ReflectionProto) qualifyMessageOptionName(pkg, scope string, fqn string) string {
	// Message options must at least include the message scope, even if the option
	// is inside that message. We do that by requiring we have at least one
	// enclosing skip in the qualified name.
	return p.qualifyElementName(pkg, scope, fqn, 1)
}

func (p *ReflectionProto) qualifyExtensionLiteralName(pkg, scope string, fqn string) string {
	// In message literals, extensions can have package name omitted but may not
	// have any other scopes omitted. We signal that via negative arg.
	return p.qualifyElementName(pkg, scope, fqn, -1)
}

func (p *ReflectionProto) qualifyName(pkg, scope string, fqn string) string {
	return p.qualifyElementName(pkg, scope, fqn, 0)
}

func (p *ReflectionProto) qualifyElementName(pkg, scope string, fqn string, required int) string {
	if p.ForceFullyQualifiedNames {
		// forcing fully-qualified names; make sure to include preceding dot
		if fqn[0] == '.' {
			return fqn
		}
		return fmt.Sprintf(".%s", fqn)
	}

	// compute relative name (so no leading dot)
	if fqn[0] == '.' {
		fqn = fqn[1:]
	}
	if required < 0 {
		scope = pkg + "."
	} else if len(scope) > 0 && scope[len(scope)-1] != '.' {
		scope = scope + "."
	}
	count := 0
	for scope != "" {
		if strings.HasPrefix(fqn, scope) && count >= required {
			return fqn[len(scope):]
		}
		if scope == pkg+"." {
			break
		}
		pos := strings.LastIndex(scope[:len(scope)-1], ".")
		scope = scope[:pos+1]
		count++
	}
	return fqn
}

func (p *ReflectionProto) typeString(fld *desc.FieldDescriptor, scope string) string {
	if fld.IsMap() {
		return fmt.Sprintf("map<%s, %s>", p.typeString(fld.GetMapKeyType(), scope), p.typeString(fld.GetMapValueType(), scope))
	}
	switch fld.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		return "sint32"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return "sint64"
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return "fixed32"
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return "fixed64"
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return "sfixed32"
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return "sfixed64"
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return "float"
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return "double"
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return "bytes"
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		return p.qualifyName(fld.GetFile().GetPackage(), scope, fld.GetEnumType().GetFullyQualifiedName())
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		return p.qualifyName(fld.GetFile().GetPackage(), scope, fld.GetMessageType().GetFullyQualifiedName())
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
		return fld.GetMessageType().GetName()
	}
	panic(fmt.Sprintf("invalid type: %v", fld.GetType()))
}

func (p *ReflectionProto) descFile(fd *desc.FileDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap) {
	opts, err := p.extractOptions(fd, fd.GetOptions(), mf)
	if err != nil {
		return
	}

	fdp := fd.AsFileDescriptorProto()
	path := make([]int32, 1)

	path[0] = internal.File_packageTag
	sourceInfo.PutIfAbsent(append(path, 0), sourceInfo.Get(path))

	path[0] = internal.File_syntaxTag
	si := sourceInfo.Get(path)
	p.descElement(false, si, 0, func() {
		syn := fdp.GetSyntax()
		if syn == "" {
			syn = "proto2"
		}
	})
	p.newLine()

	skip := map[interface{}]bool{}

	elements := elementAddrs{dsc: fd, opts: opts}
	if fdp.Package != nil {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.File_packageTag, elementIndex: 0, order: -3})
	}
	for i := range fdp.GetDependency() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.File_dependencyTag, elementIndex: i, order: -2})
	}
	elements.addrs = append(elements.addrs, optionsAsElementAddrs(internal.File_optionsTag, -1, opts)...)
	for i := range fd.GetMessageTypes() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.File_messagesTag, elementIndex: i})
	}
	for i := range fd.GetEnumTypes() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.File_enumsTag, elementIndex: i})
	}
	for i := range fd.GetServices() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.File_servicesTag, elementIndex: i})
	}
	exts := p.computeExtensions(sourceInfo, fd.GetExtensions(), []int32{internal.File_extensionsTag})
	for i, extd := range fd.GetExtensions() {
		if extd.GetType() == descriptor.FieldDescriptorProto_TYPE_GROUP {
			// we don't emit nested messages for groups since
			// they get special treatment
			skip[extd.GetMessageType()] = true
		}
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.File_extensionsTag, elementIndex: i})
	}

	p.sort(elements, sourceInfo, nil)

	pkgName := fd.GetPackage()

	for i, el := range elements.addrs {
		d := elements.at(el)

		// skip[d] will panic if d is a slice (which it could be for []option),
		// so just ignore it since we don't try to skip options
		if reflect.TypeOf(d).Kind() != reflect.Slice && skip[d] {
			// skip this element
			continue
		}

		if i > 0 {
			p.newLine()
		}

		path = []int32{el.elementType, int32(el.elementIndex)}

		switch d := d.(type) {
		case pkg:
			si := sourceInfo.Get(path)
			p.descElement(false, si, 0, func() {
			})
		case imp:
			si := sourceInfo.Get(path)
			var modifier string
			for _, idx := range fdp.PublicDependency {
				if fdp.Dependency[idx] == string(d) {
					modifier = "public "
					break
				}
			}
			if modifier == "" {
				for _, idx := range fdp.WeakDependency {
					if fdp.Dependency[idx] == string(d) {
						modifier = "weak "
						break
					}
				}
			}
			p.descElement(false, si, 0, func() {
			})
		case []option:
			p.descOptionsLong(d, sourceInfo, path, 0)
		case *desc.MessageDescriptor:
			p.descMessage(d, mf, sourceInfo, path, 0)
		case *desc.EnumDescriptor:
			p.descEnum(d, mf, sourceInfo, path, 0)
		case *desc.ServiceDescriptor:
			p.descService(d, mf, sourceInfo, path, 0)
		case *desc.FieldDescriptor:
			extDecl := exts[d]
			p.descExtensions(extDecl, exts, elements, i, mf, sourceInfo, nil, internal.File_extensionsTag, pkgName, pkgName, 0)
			// we printed all extensions in the group, so we can skip the others
			for _, fld := range extDecl.fields {
				skip[fld] = true
			}
		}
	}
}

func (p *ReflectionProto) descExtensions(exts *extensionDecl, allExts extensions, parentElements elementAddrs, startFieldIndex int, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, parentPath []int32, extTag int32, pkg, scope string, indent int) {
	path := append(parentPath, extTag)
	p.descLeadingComments(exts.sourceInfo, indent)
	p.indent(indent)
	extNameSi := sourceInfo.Get(append(path, 0, internal.Field_extendeeTag))
	p.descElementString(extNameSi, indent, p.qualifyName(pkg, scope, exts.extendee))

	if p.descTrailingComments(exts.sourceInfo, indent+1) && !p.Compact {
		// separator line between trailing comment and next element
	}

	count := len(exts.fields)
	first := true
	for idx := startFieldIndex; count > 0 && idx < len(parentElements.addrs); idx++ {
		el := parentElements.addrs[idx]
		if el.elementType != extTag {
			continue
		}
		fld := parentElements.at(el).(*desc.FieldDescriptor)
		if allExts[fld] == exts {
			if first {
				first = false
			} else {
			}
			childPath := append(path, int32(el.elementIndex))
			p.descField(fld, mf, sourceInfo, childPath, scope, indent+1)
			count--
		}
	}

	p.indent(indent)
}

func isGroup(fld *desc.FieldDescriptor) bool {
	return fld.GetType() == descriptorpb.FieldDescriptorProto_TYPE_GROUP
}

func shouldEmitLabel(fld *desc.FieldDescriptor) bool {
	return fld.IsProto3Optional() ||
		(!fld.IsMap() && fld.GetOneOf() == nil &&
			(fld.GetLabel() != descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL || !fld.GetFile().IsProto3()))
}

func labelString(lbl descriptorpb.FieldDescriptorProto_Label) string {
	switch lbl {
	case descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL:
		return "optional"
	case descriptorpb.FieldDescriptorProto_LABEL_REQUIRED:
		return "required"
	case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
		return "repeated"
	}
	panic(fmt.Sprintf("invalid label: %v", lbl))
}

func (p *ReflectionProto) descField(fld *desc.FieldDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, scope string, indent int) {
	var groupPath []int32
	var si *descriptor.SourceCodeInfo_Location

	group := isGroup(fld)

	if group {
		// compute path to group message type
		groupPath = make([]int32, len(path)-2)
		copy(groupPath, path)

		var candidates []*desc.MessageDescriptor
		var parentTag int32
		switch parent := fld.GetParent().(type) {
		case *desc.MessageDescriptor:
			// group in a message
			candidates = parent.GetNestedMessageTypes()
			parentTag = internal.Message_nestedMessagesTag
		case *desc.FileDescriptor:
			// group that is a top-level extension
			candidates = parent.GetMessageTypes()
			parentTag = internal.File_messagesTag
		}

		var groupMsgIndex int32
		for i, nmd := range candidates {
			if nmd == fld.GetMessageType() {
				// found it
				groupMsgIndex = int32(i)
				break
			}
		}
		groupPath = append(groupPath, parentTag, groupMsgIndex)

		// the group message is where the field's comments and position are stored
		si = sourceInfo.Get(groupPath)
	} else {
		si = sourceInfo.Get(path)
	}

	p.descElement(true, si, indent, func() {
		p.indent(indent)
		if shouldEmitLabel(fld) {
			locSi := sourceInfo.Get(append(path, internal.Field_labelTag))
			p.descElementString(locSi, indent, labelString(fld.GetLabel()))
		}

		if group {
		}

		typeSi := sourceInfo.Get(append(path, internal.Field_typeTag))
		p.descElementString(typeSi, indent, p.typeString(fld, scope))

		if !group {
			nameSi := sourceInfo.Get(append(path, internal.Field_nameTag))
			p.descElementString(nameSi, indent, fld.GetName())
		}

		numSi := sourceInfo.Get(append(path, internal.Field_numberTag))
		p.descElementString(numSi, indent, fmt.Sprintf("%d", fld.GetNumber()))

		opts, err := p.extractOptions(fld, fld.GetOptions(), mf)
		if err != nil {
			return
		}

		// we use negative values for "extras" keys so they can't collide
		// with legit option tags

		if !fld.GetFile().IsProto3() && fld.AsFieldDescriptorProto().DefaultValue != nil {
			defVal := fld.GetDefaultValue()
			if fld.GetEnumType() != nil {
				defVal = fld.GetEnumType().FindValueByNumber(defVal.(int32))
			}
			opts[-internal.Field_defaultTag] = []option{{name: "default", val: defVal}}
		}

		jsn := fld.AsFieldDescriptorProto().GetJsonName()
		if jsn != "" && jsn != internal.JsonName(fld.GetName()) {
			opts[-internal.Field_jsonNameTag] = []option{{name: "json_name", val: jsn}}
		}

		elements := elementAddrs{dsc: fld, opts: opts}
		elements.addrs = optionsAsElementAddrs(internal.Field_optionsTag, 0, opts)
		p.sort(elements, sourceInfo, path)
		p.descOptionElementsShort(elements, sourceInfo, path, indent)

		if group {
			p.descMessageBody(fld.GetMessageType(), mf, sourceInfo, groupPath, indent+1)

			p.indent(indent)

		} else {
		}
	})
}

func (p *ReflectionProto) descExtensionRanges(parent *desc.MessageDescriptor, ranges []*descriptorpb.DescriptorProto_ExtensionRange, maxTag int32, addrs []elementAddr, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, parentPath []int32, indent int) {
	p.indent(indent)

	var opts *descriptor.ExtensionRangeOptions
	var elPath []int32
	first := true
	for i, extr := range ranges {
		if first {
			first = false
		} else {
		}
		opts = extr.Options
		el := addrs[i]
		elPath = append(parentPath, el.elementType, int32(el.elementIndex))
		si := sourceInfo.Get(elPath)
		p.descElement(true, si, inline(indent), func() {
			if extr.GetStart() == extr.GetEnd()-1 {
			} else if extr.GetEnd()-1 == maxTag {
			} else {
			}
		})
	}
	dsc := extensionRange{owner: parent, extRange: ranges[0]}
	p.descOptionsShort(dsc, opts, mf, internal.ExtensionRange_optionsTag, sourceInfo, elPath, indent)

}

func (p *ReflectionProto) descService(sd *desc.ServiceDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	serviceInfo := &ServiceInfo{
		ServiceDescriptor: sd,
	}

	si := sourceInfo.Get(path)
	p.descElement(true, si, indent, func() {
		p.indent(indent)

		nameSi := sourceInfo.Get(append(path, internal.Service_nameTag))
		p.descElementString(nameSi, indent, sd.GetName())

		indent++

		opts, err := p.extractOptions(sd, sd.GetOptions(), mf)
		if err != nil {
			return
		}

		elements := elementAddrs{dsc: sd, opts: opts}
		elements.addrs = append(elements.addrs, optionsAsElementAddrs(internal.Service_optionsTag, -1, opts)...)
		for i := range sd.GetMethods() {
			elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Service_methodsTag, elementIndex: i})
		}

		p.sort(elements, sourceInfo, path)

		for i, el := range elements.addrs {
			if i > 0 {
			}

			childPath := append(path, el.elementType, int32(el.elementIndex))

			switch d := elements.at(el).(type) {
			case []option:
				//p.descOptionsLong(d, reg, sourceInfo, childPath, indent)
				serviceInfo.Options = append(serviceInfo.Options, d...)
			case *desc.MethodDescriptor:
				p.descMethod(d, mf, sourceInfo, childPath, indent)
				serviceInfo.MethodDescriptor = append(serviceInfo.MethodDescriptor, d)
			}
		}

		p.indent(indent - 1)
	})

	p.FileServiceInfo = append(p.FileServiceInfo, serviceInfo)
}

func (p *ReflectionProto) extractOptions(dsc desc.Descriptor, opts protov1.Message, mf *dynamic.MessageFactory) (map[int32][]option, error) {
	md, err := desc.LoadMessageDescriptorForMessage(opts)
	if err != nil {
		return nil, err
	}
	dm := mf.NewDynamicMessage(md)
	if err = dm.ConvertFrom(opts); err != nil {
		return nil, fmt.Errorf("failed convert %s to dynamic message: %v", md.GetFullyQualifiedName(), err)
	}

	pkg := dsc.GetFile().GetPackage()
	var scope string
	if _, ok := dsc.(*desc.FileDescriptor); ok {
		scope = pkg
	} else {
		scope = dsc.GetFullyQualifiedName()
	}

	options := map[int32][]option{}
	var uninterpreted []interface{}
	for _, fldset := range [][]*desc.FieldDescriptor{md.GetFields(), mf.GetExtensionRegistry().AllExtensionsForType(md.GetFullyQualifiedName())} {
		for _, fld := range fldset {
			if dm.HasField(fld) {
				val := dm.GetField(fld)
				var opts []option
				var name string
				if fld.IsExtension() {
					name = fmt.Sprintf("(%s)", p.qualifyName(pkg, scope, fld.GetFullyQualifiedName()))
				} else {
					name = fld.GetName()
				}
				switch val := val.(type) {
				case []interface{}:
					if fld.GetNumber() == internal.UninterpretedOptionsTag {
						// we handle uninterpreted options differently
						uninterpreted = val
						continue
					}

					for _, e := range val {
						if fld.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
							ev := fld.GetEnumType().FindValueByNumber(e.(int32))
							if ev == nil {
								// have to skip unknown enum values :(
								continue
							}
							e = ev
						}
						opts = append(opts, option{name: name, val: e})
					}
				case map[interface{}]interface{}:
					for k := range sortKeys(val) {
						v := val[k]
						vf := fld.GetMapValueType()
						if vf.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
							ev := vf.GetEnumType().FindValueByNumber(v.(int32))
							if ev == nil {
								// have to skip unknown enum values :(
								continue
							}
							v = ev
						}
						entry := mf.NewDynamicMessage(fld.GetMessageType())
						entry.SetFieldByNumber(1, k)
						entry.SetFieldByNumber(2, v)
						opts = append(opts, option{name: name, val: entry})
					}
				default:
					if fld.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {
						ev := fld.GetEnumType().FindValueByNumber(val.(int32))
						if ev == nil {
							// have to skip unknown enum values :(
							continue
						}
						val = ev
					}
					opts = append(opts, option{name: name, val: val})
				}
				if len(opts) > 0 {
					options[fld.GetNumber()] = opts
				}
			}
		}
	}

	// if there are uninterpreted options, add those too
	if len(uninterpreted) > 0 {
		opts := make([]option, len(uninterpreted))
		for i, u := range uninterpreted {
			var unint *descriptor.UninterpretedOption
			if un, ok := u.(*descriptor.UninterpretedOption); ok {
				unint = un
			} else {
				dm := u.(*dynamic.Message)
				unint = &descriptor.UninterpretedOption{}
				if err := dm.ConvertTo(unint); err != nil {
					return nil, err
				}
			}

			var buf bytes.Buffer
			for ni, n := range unint.Name {
				if ni > 0 {
					buf.WriteByte('.')
				}
				if n.GetIsExtension() {
					fmt.Fprintf(&buf, "(%s)", n.GetNamePart())
				} else {
					buf.WriteString(n.GetNamePart())
				}
			}

			var v interface{}
			switch {
			case unint.IdentifierValue != nil:
				v = ident(unint.GetIdentifierValue())
			case unint.StringValue != nil:
				v = string(unint.GetStringValue())
			case unint.DoubleValue != nil:
				v = unint.GetDoubleValue()
			case unint.PositiveIntValue != nil:
				v = unint.GetPositiveIntValue()
			case unint.NegativeIntValue != nil:
				v = unint.GetNegativeIntValue()
			case unint.AggregateValue != nil:
				v = ident(unint.GetAggregateValue())
			}

			opts[i] = option{name: buf.String(), val: v}
		}
		options[internal.UninterpretedOptionsTag] = opts
	}

	return options, nil
}

func (p *ReflectionProto) descMethod(mtd *desc.MethodDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	si := sourceInfo.Get(path)
	pkg := mtd.GetFile().GetPackage()
	p.descBlockElement(true, si, indent, func(trailer func(int, bool)) {
		p.indent(indent)

		nameSi := sourceInfo.Get(append(path, internal.Method_nameTag))
		p.descElementString(nameSi, indent, mtd.GetName())

		inSi := sourceInfo.Get(append(path, internal.Method_inputTag))
		inName := p.qualifyName(pkg, pkg, mtd.GetInputType().GetFullyQualifiedName())
		if mtd.IsClientStreaming() {
			inName = "stream " + inName
		}
		p.descElementString(inSi, indent, inName)

		outSi := sourceInfo.Get(append(path, internal.Method_outputTag))
		outName := p.qualifyName(pkg, pkg, mtd.GetOutputType().GetFullyQualifiedName())
		if mtd.IsServerStreaming() {
			outName = "stream " + outName
		}
		p.descElementString(outSi, indent, outName)

		opts, err := p.extractOptions(mtd, mtd.GetOptions(), mf)
		if err != nil {
			return
		}

		if len(opts) > 0 {
			indent++
			trailer(indent, true)

			elements := elementAddrs{dsc: mtd, opts: opts}
			elements.addrs = optionsAsElementAddrs(internal.Method_optionsTag, 0, opts)
			p.sort(elements, sourceInfo, path)

			for i, el := range elements.addrs {
				if i > 0 {
				}
				o := elements.at(el).([]option)
				//childPath := append(path, el.elementType, int32(el.elementIndex))
				//p.descOptionsLong(o, reg, sourceInfo, childPath, indent)
				p.MethodOptions[mtd] = append(p.MethodOptions[mtd], o...)
			}

			p.indent(indent - 1)
		} else {
			trailer(indent, false)
		}
	})
}

func (p *ReflectionProto) descEnum(ed *desc.EnumDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	si := sourceInfo.Get(path)
	p.descElement(true, si, indent, func() {
		p.indent(indent)

		nameSi := sourceInfo.Get(append(path, internal.Enum_nameTag))
		p.descElementString(nameSi, indent, ed.GetName())

		indent++
		opts, err := p.extractOptions(ed, ed.GetOptions(), mf)
		if err != nil {
			return
		}

		skip := map[interface{}]bool{}

		elements := elementAddrs{dsc: ed, opts: opts}
		elements.addrs = append(elements.addrs, optionsAsElementAddrs(internal.Enum_optionsTag, -1, opts)...)
		for i := range ed.GetValues() {
			elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Enum_valuesTag, elementIndex: i})
		}
		for i := range ed.AsEnumDescriptorProto().GetReservedRange() {
			elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Enum_reservedRangeTag, elementIndex: i})
		}
		for i := range ed.AsEnumDescriptorProto().GetReservedName() {
			elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Enum_reservedNameTag, elementIndex: i})
		}

		p.sort(elements, sourceInfo, path)

		for i, el := range elements.addrs {
			d := elements.at(el)

			// skip[d] will panic if d is a slice (which it could be for []option),
			// so just ignore it since we don't try to skip options
			if reflect.TypeOf(d).Kind() != reflect.Slice && skip[d] {
				// skip this element
				continue
			}

			if i > 0 {
				p.newLine()
			}

			childPath := append(path, el.elementType, int32(el.elementIndex))

			switch d := d.(type) {
			case []option:
				p.descOptionsLong(d, sourceInfo, childPath, indent)
			case *desc.EnumValueDescriptor:
				p.descEnumValue(d, mf, sourceInfo, childPath, indent)
			case reservedRange:
				// collapse reserved ranges into a single "reserved" block
				ranges := []reservedRange{d}
				addrs := []elementAddr{el}
				for idx := i + 1; idx < len(elements.addrs); idx++ {
					elnext := elements.addrs[idx]
					if elnext.elementType != el.elementType {
						break
					}
					rr := elements.at(elnext).(reservedRange)
					ranges = append(ranges, rr)
					addrs = append(addrs, elnext)
					skip[rr] = true
				}
				p.descReservedRanges(ranges, math.MaxInt32, addrs, sourceInfo, path, indent)
			case string: // reserved name
				// collapse reserved names into a single "reserved" block
				names := []string{d}
				addrs := []elementAddr{el}
				for idx := i + 1; idx < len(elements.addrs); idx++ {
					elnext := elements.addrs[idx]
					if elnext.elementType != el.elementType {
						break
					}
					rn := elements.at(elnext).(string)
					names = append(names, rn)
					addrs = append(addrs, elnext)
					skip[rn] = true
				}
				p.descReservedNames(names, addrs, sourceInfo, path, indent)
			}
		}

		p.indent(indent - 1)
	})
}

func (p *ReflectionProto) descReservedRanges(ranges []reservedRange, maxVal int32, addrs []elementAddr, sourceInfo internal.SourceInfoMap, parentPath []int32, indent int) {
	p.indent(indent)
	// reserved

	first := true
	for i, rr := range ranges {
		if first {
			first = false
		} else {
		}
		el := addrs[i]
		si := sourceInfo.Get(append(parentPath, el.elementType, int32(el.elementIndex)))
		p.descElement(false, si, inline(indent), func() {
			if rr.start == rr.end {
			} else if rr.end == maxVal {
			} else {
			}
		})
	}

}

func (p *ReflectionProto) descReservedNames(names []string, addrs []elementAddr, sourceInfo internal.SourceInfoMap, parentPath []int32, indent int) {
	p.indent(indent)

	first := true
	for i, name := range names {
		if first {
			first = false
		} else {
		}
		el := addrs[i]
		si := sourceInfo.Get(append(parentPath, el.elementType, int32(el.elementIndex)))
		p.descElementString(si, indent, quotedString(name))
	}

}

func (p *ReflectionProto) descEnumValue(evd *desc.EnumValueDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	si := sourceInfo.Get(path)
	p.descElement(true, si, indent, func() {
		p.indent(indent)

		nameSi := sourceInfo.Get(append(path, internal.EnumVal_nameTag))
		p.descElementString(nameSi, indent, evd.GetName())

		numSi := sourceInfo.Get(append(path, internal.EnumVal_numberTag))
		p.descElementString(numSi, indent, fmt.Sprintf("%d", evd.GetNumber()))

		p.descOptionsShort(evd, evd.GetOptions(), mf, internal.EnumVal_optionsTag, sourceInfo, path, indent)

	})
}

func (p *ReflectionProto) descMessage(md *desc.MessageDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	si := sourceInfo.Get(path)
	p.descBlockElement(true, si, indent, func(trailer func(int, bool)) {
		p.indent(indent)

		//message
		nameSi := sourceInfo.Get(append(path, internal.Message_nameTag))
		p.descElementString(nameSi, indent, md.GetName())
		// {
		trailer(indent+1, true)

		p.descMessageBody(md, mf, sourceInfo, path, indent+1)
		p.indent(indent)
		//}
	})
}

func (p *ReflectionProto) descElementString(si *descriptorpb.SourceCodeInfo_Location, indent int, str string) {
	p.descElement(false, si, inline(indent), func() {
	})
}

func (p *ReflectionProto) descBlockElement(isDecriptor bool, si *descriptorpb.SourceCodeInfo_Location, indent int, el func(trailer func(indent int, wantTrailingNewline bool))) {
	includeComments := isDecriptor || p.includeCommentType(CommentsTokens)

	if includeComments && si != nil {
		p.descLeadingComments(si, indent)
	}
	el(func(indent int, wantTrailingNewline bool) {
		if includeComments && si != nil {
			if p.descTrailingComments(si, indent) && wantTrailingNewline && !p.Compact {
				// separator line between trailing comment and next element
			}
		}
	})
	if indent >= 0 {
		// if we're not printing inline but element did not have trailing newline, add one now
	}
}

func (p *ReflectionProto) descMessageBody(md *desc.MessageDescriptor, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	opts, err := p.extractOptions(md, md.GetOptions(), mf)
	if err != nil {
		return
	}

	skip := map[interface{}]bool{}
	maxTag := internal.GetMaxTag(md.GetMessageOptions().GetMessageSetWireFormat())

	elements := elementAddrs{dsc: md, opts: opts}
	elements.addrs = append(elements.addrs, optionsAsElementAddrs(internal.Message_optionsTag, -1, opts)...)
	for i := range md.AsDescriptorProto().GetReservedRange() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_reservedRangeTag, elementIndex: i})
	}
	for i := range md.AsDescriptorProto().GetReservedName() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_reservedNameTag, elementIndex: i})
	}
	for i := range md.AsDescriptorProto().GetExtensionRange() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_extensionRangeTag, elementIndex: i})
	}
	for i, fld := range md.GetFields() {
		if fld.IsMap() || fld.GetType() == descriptor.FieldDescriptorProto_TYPE_GROUP {
			// we don't emit nested messages for map types or groups since
			// they get special treatment
			skip[fld.GetMessageType()] = true
		}
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_fieldsTag, elementIndex: i})
	}
	for i := range md.GetNestedMessageTypes() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_nestedMessagesTag, elementIndex: i})
	}
	for i := range md.GetNestedEnumTypes() {
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_enumsTag, elementIndex: i})
	}
	exts := p.computeExtensions(sourceInfo, md.GetNestedExtensions(), append(path, internal.Message_extensionsTag))
	for i, extd := range md.GetNestedExtensions() {
		if extd.GetType() == descriptor.FieldDescriptorProto_TYPE_GROUP {
			// we don't emit nested messages for groups since
			// they get special treatment
			skip[extd.GetMessageType()] = true
		}
		elements.addrs = append(elements.addrs, elementAddr{elementType: internal.Message_extensionsTag, elementIndex: i})
	}

	p.sort(elements, sourceInfo, path)

	pkg := md.GetFile().GetPackage()
	scope := md.GetFullyQualifiedName()

	for i, el := range elements.addrs {
		d := elements.at(el)

		// skip[d] will panic if d is a slice (which it could be for []option),
		// so just ignore it since we don't try to skip options
		if reflect.TypeOf(d).Kind() != reflect.Slice && skip[d] {
			// skip this element
			continue
		}

		if i > 0 {
			p.newLine()
		}

		childPath := append(path, el.elementType, int32(el.elementIndex))

		switch d := d.(type) {
		case []option:
			p.descOptionsLong(d, sourceInfo, childPath, indent)
		case *desc.FieldDescriptor:
			if d.IsExtension() {
				extDecl := exts[d]
				p.descExtensions(extDecl, exts, elements, i, mf, sourceInfo, path, internal.Message_extensionsTag, pkg, scope, indent)
				// we printed all extensions in the group, so we can skip the others
				for _, fld := range extDecl.fields {
					skip[fld] = true
				}
			} else {
				ood := d.GetOneOf()
				if ood == nil || ood.IsSynthetic() {
					p.descField(d, mf, sourceInfo, childPath, scope, indent)
				} else {
					// print the one-of, including all of its fields
					p.descOneOf(ood, elements, i, mf, sourceInfo, path, indent, d.AsFieldDescriptorProto().GetOneofIndex())
					for _, fld := range ood.GetChoices() {
						skip[fld] = true
					}
				}
			}
		case *desc.MessageDescriptor:
			p.descMessage(d, mf, sourceInfo, childPath, indent)
		case *desc.EnumDescriptor:
			p.descEnum(d, mf, sourceInfo, childPath, indent)
		case *descriptor.DescriptorProto_ExtensionRange:
			// collapse ranges into a single "extensions" block
			ranges := []*descriptor.DescriptorProto_ExtensionRange{d}
			addrs := []elementAddr{el}
			for idx := i + 1; idx < len(elements.addrs); idx++ {
				elnext := elements.addrs[idx]
				if elnext.elementType != el.elementType {
					break
				}
				extr := elements.at(elnext).(*descriptor.DescriptorProto_ExtensionRange)
				if !areEqual(d.Options, extr.Options, mf) {
					break
				}
				ranges = append(ranges, extr)
				addrs = append(addrs, elnext)
				skip[extr] = true
			}
			p.descExtensionRanges(md, ranges, maxTag, addrs, mf, sourceInfo, path, indent)
		case reservedRange:
			// collapse reserved ranges into a single "reserved" block
			ranges := []reservedRange{d}
			addrs := []elementAddr{el}
			for idx := i + 1; idx < len(elements.addrs); idx++ {
				elnext := elements.addrs[idx]
				if elnext.elementType != el.elementType {
					break
				}
				rr := elements.at(elnext).(reservedRange)
				ranges = append(ranges, rr)
				addrs = append(addrs, elnext)
				skip[rr] = true
			}
			p.descReservedRanges(ranges, maxTag, addrs, sourceInfo, path, indent)
		case string: // reserved name
			// collapse reserved names into a single "reserved" block
			names := []string{d}
			addrs := []elementAddr{el}
			for idx := i + 1; idx < len(elements.addrs); idx++ {
				elnext := elements.addrs[idx]
				if elnext.elementType != el.elementType {
					break
				}
				rn := elements.at(elnext).(string)
				names = append(names, rn)
				addrs = append(addrs, elnext)
				skip[rn] = true
			}
			p.descReservedNames(names, addrs, sourceInfo, path, indent)
		}
	}
}

func areEqual(a, b protov1.Message, mf *dynamic.MessageFactory) bool {
	// proto.Equal doesn't handle unknown extensions very well :(
	// so we convert to a dynamic message (which should know about all extensions via
	// extension registry) and then compare
	return dynamic.MessagesEqual(asDynamicIfPossible(a, mf), asDynamicIfPossible(b, mf))
}

func asDynamicIfPossible(msg protov1.Message, mf *dynamic.MessageFactory) protov1.Message {
	if dm, ok := msg.(*dynamic.Message); ok {
		return dm
	} else {
		md, err := desc.LoadMessageDescriptorForMessage(msg)
		if err == nil {
			dm := mf.NewDynamicMessage(md)
			if dm.ConvertFrom(msg) == nil {
				return dm
			}
		}
	}
	return msg
}

func (p *ReflectionProto) descOneOf(ood *desc.OneOfDescriptor, parentElements elementAddrs, startFieldIndex int, mf *dynamic.MessageFactory, sourceInfo internal.SourceInfoMap, parentPath []int32, indent int, ooIndex int32) {
	oopath := append(parentPath, internal.Message_oneOfsTag, ooIndex)
	oosi := sourceInfo.Get(oopath)
	p.descElement(true, oosi, indent, func() {
		p.indent(indent)
		extNameSi := sourceInfo.Get(append(oopath, internal.OneOf_nameTag))
		p.descElementString(extNameSi, indent, ood.GetName())

		indent++
		opts, err := p.extractOptions(ood, ood.GetOptions(), mf)
		if err != nil {
			return
		}

		elements := elementAddrs{dsc: ood, opts: opts}
		elements.addrs = append(elements.addrs, optionsAsElementAddrs(internal.OneOf_optionsTag, -1, opts)...)

		count := len(ood.GetChoices())
		for idx := startFieldIndex; count > 0 && idx < len(parentElements.addrs); idx++ {
			el := parentElements.addrs[idx]
			if el.elementType != internal.Message_fieldsTag {
				continue
			}
			if parentElements.at(el).(*desc.FieldDescriptor).GetOneOf() == ood {
				// negative tag indicates that this element is actually a sibling, not a child
				elements.addrs = append(elements.addrs, elementAddr{elementType: -internal.Message_fieldsTag, elementIndex: el.elementIndex})
				count--
			}
		}

		// the fields are already sorted, but we have to re-sort in order to
		// interleave the options (in the event that we are using file location
		// order and the option locations are interleaved with the fields)
		p.sort(elements, sourceInfo, oopath)
		scope := ood.GetOwner().GetFullyQualifiedName()

		for i, el := range elements.addrs {
			if i > 0 {
				p.newLine()
			}

			switch d := elements.at(el).(type) {
			case []option:
				childPath := append(oopath, el.elementType, int32(el.elementIndex))
				p.descOptionsLong(d, sourceInfo, childPath, indent)
			case *desc.FieldDescriptor:
				childPath := append(parentPath, -el.elementType, int32(el.elementIndex))
				p.descField(d, mf, sourceInfo, childPath, scope, indent)
			}
		}

		p.indent(indent - 1)
	})
}

func (p *ReflectionProto) descElement(isDecriptor bool, si *descriptorpb.SourceCodeInfo_Location, indent int, el func()) {
	includeComments := isDecriptor || p.includeCommentType(CommentsTokens)

	if includeComments && si != nil {
		p.descLeadingComments(si, indent)
	}
	el()
	if includeComments && si != nil {
		p.descTrailingComments(si, indent)
	}
}

func (p *ReflectionProto) includeCommentType(c CommentType) bool {
	return (p.OmitComments & c) == 0
}

func (p *ReflectionProto) descLeadingComments(si *descriptorpb.SourceCodeInfo_Location, indent int) bool {
	endsInNewLine := false

	if p.includeCommentType(CommentsDetached) {
		for _, c := range si.GetLeadingDetachedComments() {
			if p.descComment(c, indent, true) {
				// if comment ended in newline, add another newline to separate
				// this comment from the next
				p.newLine()
				endsInNewLine = true
			} else if indent < 0 {
				// comment did not end in newline and we are trying to inline?
				// just add a space to separate this comment from what follows
				endsInNewLine = false
			} else {
				// comment did not end in newline and we are *not* trying to inline?
				// add newline to end of comment and add another to separate this
				// comment from what follows
				p.newLine()
				endsInNewLine = true
			}
		}
	}

	if p.includeCommentType(CommentsLeading) && si.GetLeadingComments() != "" {
		endsInNewLine = p.descComment(si.GetLeadingComments(), indent, true)
		if !endsInNewLine {
			if indent >= 0 {
				// leading comment didn't end with newline but needs one
				// (because we're *not* inlining)
				endsInNewLine = true
			} else {
				// space between comment and following element when inlined
			}
		}
	}

	return endsInNewLine
}

func (p *ReflectionProto) descTrailingComments(si *descriptorpb.SourceCodeInfo_Location, indent int) bool {
	if p.includeCommentType(CommentsTrailing) && si.GetTrailingComments() != "" {
		if !p.descComment(si.GetTrailingComments(), indent, p.TrailingCommentsOnSeparateLine) && indent >= 0 {
			// trailing comment didn't end with newline but needs one
			// (because we're *not* inlining)
		} else if indent < 0 {
		}
		return true
	}

	return false
}

func (p *ReflectionProto) newLine() {
	if !p.Compact {
	}
}

func (p *ReflectionProto) descComment(comments string, indent int, forceNextLine bool) bool {
	if comments == "" {
		return false
	}

	var multiLine bool
	if indent < 0 {
		// use multi-line style when inlining
		multiLine = true
	} else {
		multiLine = p.PreferMultiLineStyleComments
	}
	if multiLine && strings.Contains(comments, "*/") {
		// can't emit '*/' in a multi-line style comment
		multiLine = false
	}

	lines := strings.Split(comments, "\n")

	// first, remove leading and trailing blank lines
	if lines[0] == "" {
		lines = lines[1:]
	}
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return false
	}

	if indent >= 0 {
		// last element did not have trailing newline, so we
		// either need to tack on newline or, if comment is
		// just one line, inline it on the end
		if forceNextLine || len(lines) > 1 {
		} else {
			indent = inline(indent)
		}
	}

	if len(lines) == 1 && multiLine {
		p.indent(indent)
		line := lines[0]
		if line[0] == ' ' && line[len(line)-1] != ' ' {
			// add trailing space for symmetry
			line += " "
		}
		p.Comment = append(p.Comment, fmt.Sprintf("/*%s*/", line))
		if indent >= 0 {
			return true
		}
		return false
	}

	if multiLine {
		// multi-line style comments that actually span multiple lines
		// get a blank line before and after so that comment renders nicely
		lines = append(lines, "", "")
		copy(lines[1:], lines)
		lines[0] = ""
	}

	var sb strings.Builder
	for i, l := range lines {
		if l != "" && !strings.HasPrefix(l, " ") {
			l = " " + l
		}
		p.maybeIndent(indent, i > 0)
		if multiLine {
			if i == 0 {
				// first line
				sb.WriteString(fmt.Sprintf("/*%s\n", strings.TrimRight(l, " \t")))
			} else if i == len(lines)-1 {
				// last line
				if strings.TrimSpace(l) == "" {
					sb.WriteString(fmt.Sprint(" */"))
				} else {
					sb.WriteString(fmt.Sprintf(" *%s*/", l))
				}
				if indent >= 0 {
					sb.WriteString(fmt.Sprintln())
				}
			} else {
				sb.WriteString(fmt.Sprintf(" *%s\n", strings.TrimRight(l, " \t")))
			}
		} else {
			p.Comment = append(p.Comment, fmt.Sprintf("//%s\n", strings.TrimRight(l, " \t")))
		}
	}

	if sb.Len() > 0 {
		p.Comment = append(p.Comment, sb.String())
	}

	// single-line comments always end in newline; multi-line comments only
	// end in newline for non-negative (e.g. non-inlined) indentation
	return !multiLine || indent >= 0
}

func (p *ReflectionProto) indent(indent int) {
	for i := 0; i < indent; i++ {
	}
}

func (p *ReflectionProto) maybeIndent(indent int, requireIndent bool) {
	if indent < 0 && requireIndent {
		p.indent(-indent)
	} else {
		p.indent(indent)
	}
}

func inline(indent int) int {
	if indent < 0 {
		// already inlined
		return indent
	}
	// negative indent means inline; indent 2 stops further in case value wraps
	return -indent - 2
}

func (p *ReflectionProto) descOptionsLong(opts []option, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	p.descOptions(opts, indent,
		func(i int32) *descriptor.SourceCodeInfo_Location {
			return sourceInfo.Get(append(path, i))
		},
		func(indent int, opt option) {
			p.indent(indent)
			p.descOption(opt.name, opt.val, indent)
		})
}

func (p *ReflectionProto) descOptionsShort(dsc interface{}, optsMsg protov1.Message, mf *dynamic.MessageFactory, optsTag int32, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	d, ok := dsc.(desc.Descriptor)
	if !ok {
		d = dsc.(extensionRange).owner
	}
	opts, err := p.extractOptions(d, optsMsg, mf)
	if err != nil {
		return
	}

	elements := elementAddrs{dsc: dsc, opts: opts}
	elements.addrs = optionsAsElementAddrs(optsTag, 0, opts)
	p.sort(elements, sourceInfo, path)
	p.descOptionElementsShort(elements, sourceInfo, path, indent)
}

func (p *ReflectionProto) descOptionElementsShort(addrs elementAddrs, sourceInfo internal.SourceInfoMap, path []int32, indent int) {
	if len(addrs.addrs) == 0 {
		return
	}
	first := true
	for _, addr := range addrs.addrs {
		opts := addrs.at(addr).([]option)
		var childPath []int32
		if addr.elementIndex < 0 {
			// pseudo-option
			childPath = append(path, int32(-addr.elementIndex))
		} else {
			childPath = append(path, addr.elementType, int32(addr.elementIndex))
		}
		p.descOptions(opts, inline(indent),
			func(i int32) *descriptor.SourceCodeInfo_Location {
				p := childPath
				if addr.elementIndex >= 0 {
					p = append(p, i)
				}
				return sourceInfo.Get(p)
			},
			func(indent int, opt option) {
				if first {
					first = false
				} else {
				}
				p.descOption(opt.name, opt.val, indent)
			})
	}
}

func (p *ReflectionProto) descOptions(opts []option, indent int, siFetch func(i int32) *descriptor.SourceCodeInfo_Location, fn func(indent int, opt option)) {
	for i, opt := range opts {
		si := siFetch(int32(i))
		p.descElement(false, si, indent, func() {
			fn(indent, opt)
		})
	}
}

func (p *ReflectionProto) descOption(name string, optVal interface{}, indent int) {

	switch optVal := optVal.(type) {
	case int32, uint32, int64, uint64:
	case float32, float64:
	case string:
	case []byte:
	case bool:
	case ident:
	case *desc.EnumValueDescriptor:
	case proto.Message:
		// TODO: if value is too long, marshal to text format with indentation to
		// make output prettier (also requires correctly indenting subsequent lines)

		// TODO: alternate approach so we can apply p.ForceFullyQualifiedNames
		// inside the resulting value?

	default:
		panic(fmt.Sprintf("unknown type of value %T for field %s", optVal, name))
	}
}

func (p *ReflectionProto) descMessageLiteralToBufferMaybeCompact(msg protoreflect.Message, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
	if indent >= 0 {
		// first see if the message is compact enough to fit on one line
		p.descMessageLiteralCompact(msg, mf, pkg, scope)
	}
	p.descMessageLiteralToBuffer(msg, mf, pkg, scope, threshold, indent)
}

func (p *ReflectionProto) descMessageLiteralCompact(msg protoreflect.Message, mf *dynamic.MessageFactory, pkg, scope string) {
	p.descMessageLiteralToBuffer(msg, mf, pkg, scope, 0, -1)
}

func (p *ReflectionProto) descMessageLiteral(msg protoreflect.Message, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
	p.descMessageLiteralToBuffer(msg, mf, pkg, scope, threshold, indent)
}

var (
	anyTypeName = (&anypb.Any{}).ProtoReflect().Descriptor().FullName()
)

const (
	anyTypeUrlTag = 1
	anyValueTag   = 2
)

func (p *ReflectionProto) descMessageLiteralToBuffer(msg protoreflect.Message, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
	if p.maybePrintAnyMessageToBuffer(msg, mf, pkg, scope, threshold, indent) {
		return
	}

	if indent >= 0 {
		indent++
	}

	type fieldVal struct {
		fld protoreflect.FieldDescriptor
		val protoreflect.Value
	}
	var fields []fieldVal
	msg.Range(func(fld protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		fields = append(fields, fieldVal{fld: fld, val: val})
		return true
	})
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].fld.Number() < fields[j].fld.Number()
	})

	for i, fldVal := range fields {
		fld, val := fldVal.fld, fldVal.val
		if i > 0 {
		}
		p.maybeNewline(indent)
		if fld.IsExtension() {
		} else {
		}
		switch {
		case fld.IsList():
			p.descArrayLiteralToBufferMaybeCompact(fld, val.List(), mf, pkg, scope, threshold, indent)
		case fld.IsMap():
			p.descMapLiteralToBufferMaybeCompact(fld, val.Map(), mf, pkg, scope, threshold, indent)
		case fld.Kind() == protoreflect.MessageKind || fld.Kind() == protoreflect.GroupKind:
			p.descMessageLiteralToBufferMaybeCompact(val.Message(), mf, pkg, scope, threshold, indent)
		default:
			p.descValueLiteralToBuffer(fld, val.Interface())
		}
	}

	if indent >= 0 {
		indent--
	}
	p.maybeNewline(indent)
}

func (p *ReflectionProto) maybePrintAnyMessageToBuffer(msg protoreflect.Message, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) bool {
	/*md := msg.Descriptor()
	if md.FullName() != anyTypeName {
		return false
	}
	typeUrlFld := md.Fields().ByNumber(anyTypeUrlTag)
	if typeUrlFld == nil || typeUrlFld.Kind() != protoreflect.StringKind || typeUrlFld.Cardinality() == protoreflect.Repeated {
		return false
	}
	valueFld := md.Fields().ByNumber(anyValueTag)
	if valueFld == nil || valueFld.Kind() != protoreflect.BytesKind || valueFld.Cardinality() == protoreflect.Repeated {
		return false
	}
	typeUrl := msg.Get(typeUrlFld).String()
	if typeUrl == "" {
		return false
	}
	mt, err := mf.FindMessageByURL(typeUrl)
	if err != nil {
		return false
	}
	valueMsg := mt.New()
	valueBytes := msg.Get(valueFld).Bytes()
	if err := (proto.UnmarshalOptions{Resolver: res}).Unmarshal(valueBytes, valueMsg.Interface()); err != nil {
		return false
	}

	if indent >= 0 {
		indent++
	}
	p.maybeNewline(indent)

	p.descMessageLiteralToBufferMaybeCompact(valueMsg, mf, pkg, scope, threshold, indent)

	if indent >= 0 {
		indent--
	}
	p.maybeNewline(indent)*/

	return true
}

func (p *ReflectionProto) descValueLiteralToBuffer(fld protoreflect.FieldDescriptor, value interface{}) {
	switch val := value.(type) {
	case protoreflect.EnumNumber:
		ev := fld.Enum().Values().ByNumber(val)
		if ev == nil {
		} else {
		}
	case string:
	case []byte:
	case int32, uint32, int64, uint64:
	case float32, float64:
	default:
	}
}

func (p *ReflectionProto) maybeNewline(indent int) {
	if indent < 0 {
		// compact form
		return
	}
	p.indent(indent)
}

func (p *ReflectionProto) descArrayLiteralToBufferMaybeCompact(fld protoreflect.FieldDescriptor, val protoreflect.List, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
	if indent >= 0 {
		// first see if the array is compact enough to fit on one line
		p.descArrayLiteralCompact(fld, val, mf, pkg, scope)
	}
	p.descArrayLiteralToBuffer(fld, val, mf, pkg, scope, threshold, indent)
}

func (p *ReflectionProto) descArrayLiteralCompact(fld protoreflect.FieldDescriptor, val protoreflect.List, mf *dynamic.MessageFactory, pkg, scope string) {
	p.descArrayLiteralToBuffer(fld, val, mf, pkg, scope, 0, -1)
}

func (p *ReflectionProto) descArrayLiteralToBuffer(fld protoreflect.FieldDescriptor, val protoreflect.List, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
	if indent >= 0 {
		indent++
	}

	for i := 0; i < val.Len(); i++ {
		if i > 0 {
		}
		p.maybeNewline(indent)
		if fld.Kind() == protoreflect.MessageKind || fld.Kind() == protoreflect.GroupKind {
			p.descMessageLiteralToBufferMaybeCompact(val.Get(i).Message(), mf, pkg, scope, threshold, indent)
		} else {
			p.descValueLiteralToBuffer(fld, val.Get(i).Interface())
		}
	}

	if indent >= 0 {
		indent--
	}
	p.maybeNewline(indent)
}

func (p *ReflectionProto) descMapLiteralToBufferMaybeCompact(fld protoreflect.FieldDescriptor, val protoreflect.Map, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
	if indent >= 0 {
		// first see if the map is compact enough to fit on one line
		p.descMapLiteralCompact(fld, val, mf, pkg, scope)
	}
	p.descMapLiteralToBuffer(fld, val, mf, pkg, scope, threshold, indent)
}

func (p *ReflectionProto) descMapLiteralCompact(fld protoreflect.FieldDescriptor, val protoreflect.Map, mf *dynamic.MessageFactory, pkg, scope string) {
	p.descMapLiteralToBuffer(fld, val, mf, pkg, scope, 0, -1)
}

func (p *ReflectionProto) descMapLiteralToBuffer(fld protoreflect.FieldDescriptor, val protoreflect.Map, mf *dynamic.MessageFactory, pkg, scope string, threshold, indent int) {
}

type mapAsList struct {
	m              protoreflect.Map
	entry          protoreflect.MessageType
	keyFld, valFld protoreflect.FieldDescriptor
	keys           []protoreflect.MapKey
}

func (m *mapAsList) Len() int {
	return len(m.keys)
}

func (m *mapAsList) Get(i int) protoreflect.Value {
	msg := m.entry.New()
	key := m.keys[i]
	msg.Set(m.keyFld, key.Value())
	val := m.m.Get(key)
	msg.Set(m.valFld, val)
	return protoreflect.ValueOfMessage(msg)
}

func (m *mapAsList) Set(_i int, _ protoreflect.Value) {
	panic("Set is not implemented")
}

func (m *mapAsList) Append(_ protoreflect.Value) {
	panic("Append is not implemented")
}

func (m *mapAsList) AppendMutable() protoreflect.Value {
	panic("AppendMutable is not implemented")
}

func (m *mapAsList) Truncate(_ int) {
	panic("Truncate is not implemented")
}

func (m *mapAsList) NewElement() protoreflect.Value {
	panic("NewElement is not implemented")
}

func (m *mapAsList) IsValid() bool {
	return true
}
