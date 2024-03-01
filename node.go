package main

import (
	"github.com/yourselfhosted/gomark/ast"
)

type Node struct {
	Type ast.NodeType `json:"type"`
	Node BaseNode     `json:"node"`
}

type BaseNode interface{}

type LineBreakNode struct{}

type ParagraphNode struct {
	Children []*Node `json:"children"`
}

type CodeBlockNode struct {
	Language string `json:"language"`
	Content  string `json:"content"`
}

type HeadingNode struct {
	Level    int     `json:"level"`
	Children []*Node `json:"children"`
}

type HorizontalRuleNode struct {
	Symbol string `json:"symbol"`
}

type BlockquoteNode struct {
	Children []*Node `json:"children"`
}

type OrderedListNode struct {
	Number   string  `json:"number"`
	Indent   int     `json:"indent"`
	Children []*Node `json:"children"`
}

type UnorderedListNode struct {
	Symbol   string  `json:"symbol"`
	Indent   int     `json:"indent"`
	Children []*Node `json:"children"`
}

type TaskListNode struct {
	Symbol   string  `json:"symbol"`
	Indent   int     `json:"indent"`
	Complete bool    `json:"complete"`
	Children []*Node `json:"children"`
}

type MathBlockNode struct {
	Content string `json:"content"`
}

type TableRowNode struct {
	Cells []string `json:"cells"`
}

type TableNode struct {
	Header    []string       `json:"header"`
	Delimiter []string       `json:"delimiter"`
	Rows      []TableRowNode `json:"rows"`
}

type EmbeddedContentNode struct {
	ResourceName string `json:"resourceName"`
	Params       string `json:"params"`
}

type TextNode struct {
	Content string `json:"content"`
}

type BoldNode struct {
	Symbol   string  `json:"symbol"`
	Children []*Node `json:"children"`
}

type ItalicNode struct {
	Symbol  string `json:"symbol"`
	Content string `json:"content"`
}

type BoldItalicNode struct {
	Symbol  string `json:"symbol"`
	Content string `json:"content"`
}

type CodeNode struct {
	Content string `json:"content"`
}

type ImageNode struct {
	AltText string `json:"altText"`
	Url     string `json:"url"`
}

type LinkNode struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type AutoLinkNode struct {
	Url       string `json:"url"`
	IsRawText bool   `json:"isRawText"`
}

type TagNode struct {
	Content string `json:"content"`
}

type StrikethroughNode struct {
	Content string `json:"content"`
}

type EscapingCharacterNode struct {
	Symbol string `json:"symbol"`
}

type MathNode struct {
	Content string `json:"content"`
}

type HighlightNode struct {
	Content string `json:"content"`
}

type SubscriptNode struct {
	Content string `json:"content"`
}

type SuperscriptNode struct {
	Content string `json:"content"`
}

type ReferencedContentNode struct {
	ResourceName string `json:"resourceName"`
	Params       string `json:"params"`
}

type SpoilerNode struct {
	Content string `json:"content"`
}

func convertFromASTNode(node ast.Node) *Node {
	n := &Node{
		Type: node.Type(),
	}

	switch node.(type) {
	case *ast.LineBreak:
		n.Node = LineBreakNode{}
	case *ast.Paragraph:
		children := convertFromASTNodes(node.(*ast.Paragraph).Children)
		n.Node = ParagraphNode{Children: children}
	case *ast.CodeBlock:
		n.Node = CodeBlockNode{Language: node.(*ast.CodeBlock).Language, Content: node.(*ast.CodeBlock).Content}
	case *ast.Heading:
		children := convertFromASTNodes(node.(*ast.Heading).Children)
		n.Node = HeadingNode{Level: node.(*ast.Heading).Level, Children: children}
	case *ast.HorizontalRule:
		n.Node = HorizontalRuleNode{Symbol: node.(*ast.HorizontalRule).Symbol}
	case *ast.Blockquote:
		children := convertFromASTNodes(node.(*ast.Blockquote).Children)
		n.Node = BlockquoteNode{Children: children}
	case *ast.OrderedList:
		children := convertFromASTNodes(node.(*ast.OrderedList).Children)
		n.Node = OrderedListNode{Number: node.(*ast.OrderedList).Number, Indent: node.(*ast.OrderedList).Indent, Children: children}
	case *ast.UnorderedList:
		children := convertFromASTNodes(node.(*ast.UnorderedList).Children)
		n.Node = UnorderedListNode{Symbol: node.(*ast.UnorderedList).Symbol, Indent: node.(*ast.UnorderedList).Indent, Children: children}
	case *ast.TaskList:
		children := convertFromASTNodes(node.(*ast.TaskList).Children)
		n.Node = TaskListNode{Symbol: node.(*ast.TaskList).Symbol, Indent: node.(*ast.TaskList).Indent, Complete: node.(*ast.TaskList).Complete, Children: children}
	case *ast.MathBlock:
		n.Node = MathBlockNode{Content: node.(*ast.MathBlock).Content}
	case *ast.Table:
		rows := []TableRowNode{}
		for _, row := range node.(*ast.Table).Rows {
			rows = append(rows, TableRowNode{Cells: row})
		}
		n.Node = TableNode{Header: node.(*ast.Table).Header, Delimiter: node.(*ast.Table).Delimiter, Rows: rows}
	case *ast.EmbeddedContent:
		n.Node = EmbeddedContentNode{ResourceName: node.(*ast.EmbeddedContent).ResourceName, Params: node.(*ast.EmbeddedContent).Params}
	case *ast.Text:
		n.Node = TextNode{Content: node.(*ast.Text).Content}
	case *ast.Bold:
		children := convertFromASTNodes(node.(*ast.Bold).Children)
		n.Node = BoldNode{Symbol: node.(*ast.Bold).Symbol, Children: children}
	case *ast.Italic:
		n.Node = ItalicNode{Symbol: node.(*ast.Italic).Symbol, Content: node.(*ast.Italic).Content}
	case *ast.BoldItalic:
		n.Node = BoldItalicNode{Symbol: node.(*ast.BoldItalic).Symbol, Content: node.(*ast.BoldItalic).Content}
	case *ast.Code:
		n.Node = CodeNode{Content: node.(*ast.Code).Content}
	case *ast.Image:
		n.Node = ImageNode{AltText: node.(*ast.Image).AltText, Url: node.(*ast.Image).URL}
	case *ast.Link:
		n.Node = LinkNode{Text: node.(*ast.Link).Text, Url: node.(*ast.Link).URL}
	case *ast.AutoLink:
		n.Node = AutoLinkNode{Url: node.(*ast.AutoLink).URL, IsRawText: node.(*ast.AutoLink).IsRawText}
	case *ast.Tag:
		n.Node = TagNode{Content: node.(*ast.Tag).Content}
	case *ast.Strikethrough:
		n.Node = StrikethroughNode{Content: node.(*ast.Strikethrough).Content}
	case *ast.EscapingCharacter:
		n.Node = EscapingCharacterNode{Symbol: node.(*ast.EscapingCharacter).Symbol}
	case *ast.Math:
		n.Node = MathNode{Content: node.(*ast.Math).Content}
	case *ast.Highlight:
		n.Node = HighlightNode{Content: node.(*ast.Highlight).Content}
	case *ast.Subscript:
		n.Node = SubscriptNode{Content: node.(*ast.Subscript).Content}
	case *ast.Superscript:
		n.Node = SuperscriptNode{Content: node.(*ast.Superscript).Content}
	case *ast.ReferencedContent:
		n.Node = ReferencedContentNode{ResourceName: node.(*ast.ReferencedContent).ResourceName, Params: node.(*ast.ReferencedContent).Params}
	case *ast.Spoiler:
		n.Node = SpoilerNode{Content: node.(*ast.Spoiler).Content}
	default:
		n.Node = TextNode{Content: ""}
	}
	return n
}

func convertFromASTNodes(rawNodes []ast.Node) []*Node {
	nodes := []*Node{}
	for _, node := range rawNodes {
		nodes = append(nodes, convertFromASTNode(node))
	}
	return nodes
}

func convertToASTNode(node *Node) ast.Node {
	switch n := node.Node.(type) {
	case LineBreakNode:
		return &ast.LineBreak{}
	case ParagraphNode:
		children := convertToASTNodes(n.Children)
		return &ast.Paragraph{Children: children}
	case CodeBlockNode:
		return &ast.CodeBlock{Language: n.Language, Content: n.Content}
	case HeadingNode:
		children := convertToASTNodes(n.Children)
		return &ast.Heading{Level: n.Level, Children: children}
	case HorizontalRuleNode:
		return &ast.HorizontalRule{Symbol: n.Symbol}
	case BlockquoteNode:
		children := convertToASTNodes(n.Children)
		return &ast.Blockquote{Children: children}
	case OrderedListNode:
		children := convertToASTNodes(n.Children)
		return &ast.OrderedList{Number: n.Number, Indent: n.Indent, Children: children}
	case UnorderedListNode:
		children := convertToASTNodes(n.Children)
		return &ast.UnorderedList{Symbol: n.Symbol, Indent: n.Indent, Children: children}
	case TaskListNode:
		children := convertToASTNodes(n.Children)
		return &ast.TaskList{Symbol: n.Symbol, Indent: n.Indent, Complete: n.Complete, Children: children}
	case MathBlockNode:
		return &ast.MathBlock{Content: n.Content}
	case TableNode:
		rows := [][]string{}
		for _, row := range n.Rows {
			rows = append(rows, row.Cells)
		}
		return &ast.Table{Header: n.Header, Delimiter: n.Delimiter, Rows: rows}
	case EmbeddedContentNode:
		return &ast.EmbeddedContent{ResourceName: n.ResourceName, Params: n.Params}
	case TextNode:
		return &ast.Text{Content: n.Content}
	case BoldNode:
		children := convertToASTNodes(n.Children)
		return &ast.Bold{Symbol: n.Symbol, Children: children}
	case ItalicNode:
		return &ast.Italic{Symbol: n.Symbol, Content: n.Content}
	case BoldItalicNode:
		return &ast.BoldItalic{Symbol: n.Symbol, Content: n.Content}
	case CodeNode:
		return &ast.Code{Content: n.Content}
	case ImageNode:
		return &ast.Image{AltText: n.AltText, URL: n.Url}
	case LinkNode:
		return &ast.Link{Text: n.Text, URL: n.Url}
	case AutoLinkNode:
		return &ast.AutoLink{URL: n.Url, IsRawText: n.IsRawText}
	case TagNode:
		return &ast.Tag{Content: n.Content}
	case StrikethroughNode:
		return &ast.Strikethrough{Content: n.Content}
	case EscapingCharacterNode:
		return &ast.EscapingCharacter{Symbol: n.Symbol}
	case MathNode:
		return &ast.Math{Content: n.Content}
	case HighlightNode:
		return &ast.Highlight{Content: n.Content}
	case SubscriptNode:
		return &ast.Subscript{Content: n.Content}
	case SuperscriptNode:
		return &ast.Superscript{Content: n.Content}
	case ReferencedContentNode:
		return &ast.ReferencedContent{ResourceName: n.ResourceName, Params: n.Params}
	case SpoilerNode:
		return &ast.Spoiler{Content: n.Content}
	default:
		return &ast.Text{}
	}
}

func convertToASTNodes(rawNodes []*Node) []ast.Node {
	nodes := []ast.Node{}
	for _, node := range rawNodes {
		nodes = append(nodes, convertToASTNode(node))
	}
	return nodes
}
