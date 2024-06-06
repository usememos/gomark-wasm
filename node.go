package main

import (
	"encoding/json"

	"github.com/usememos/gomark/ast"
)

type Node struct {
	Type  ast.NodeType `json:"type"`
	Value BaseNode     `json:"value"`
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
		n.Value = LineBreakNode{}
	case *ast.Paragraph:
		children := convertFromASTNodes(node.(*ast.Paragraph).Children)
		n.Value = ParagraphNode{Children: children}
	case *ast.CodeBlock:
		n.Value = CodeBlockNode{Language: node.(*ast.CodeBlock).Language, Content: node.(*ast.CodeBlock).Content}
	case *ast.Heading:
		children := convertFromASTNodes(node.(*ast.Heading).Children)
		n.Value = HeadingNode{Level: node.(*ast.Heading).Level, Children: children}
	case *ast.HorizontalRule:
		n.Value = HorizontalRuleNode{Symbol: node.(*ast.HorizontalRule).Symbol}
	case *ast.Blockquote:
		children := convertFromASTNodes(node.(*ast.Blockquote).Children)
		n.Value = BlockquoteNode{Children: children}
	case *ast.OrderedList:
		children := convertFromASTNodes(node.(*ast.OrderedList).Children)
		n.Value = OrderedListNode{Number: node.(*ast.OrderedList).Number, Indent: node.(*ast.OrderedList).Indent, Children: children}
	case *ast.UnorderedList:
		children := convertFromASTNodes(node.(*ast.UnorderedList).Children)
		n.Value = UnorderedListNode{Symbol: node.(*ast.UnorderedList).Symbol, Indent: node.(*ast.UnorderedList).Indent, Children: children}
	case *ast.TaskList:
		children := convertFromASTNodes(node.(*ast.TaskList).Children)
		n.Value = TaskListNode{Symbol: node.(*ast.TaskList).Symbol, Indent: node.(*ast.TaskList).Indent, Complete: node.(*ast.TaskList).Complete, Children: children}
	case *ast.MathBlock:
		n.Value = MathBlockNode{Content: node.(*ast.MathBlock).Content}
	case *ast.Table:
		rows := []TableRowNode{}
		for _, row := range node.(*ast.Table).Rows {
			rows = append(rows, TableRowNode{Cells: row})
		}
		n.Value = TableNode{Header: node.(*ast.Table).Header, Delimiter: node.(*ast.Table).Delimiter, Rows: rows}
	case *ast.EmbeddedContent:
		n.Value = EmbeddedContentNode{ResourceName: node.(*ast.EmbeddedContent).ResourceName, Params: node.(*ast.EmbeddedContent).Params}
	case *ast.Text:
		n.Value = TextNode{Content: node.(*ast.Text).Content}
	case *ast.Bold:
		children := convertFromASTNodes(node.(*ast.Bold).Children)
		n.Value = BoldNode{Symbol: node.(*ast.Bold).Symbol, Children: children}
	case *ast.Italic:
		n.Value = ItalicNode{Symbol: node.(*ast.Italic).Symbol, Content: node.(*ast.Italic).Content}
	case *ast.BoldItalic:
		n.Value = BoldItalicNode{Symbol: node.(*ast.BoldItalic).Symbol, Content: node.(*ast.BoldItalic).Content}
	case *ast.Code:
		n.Value = CodeNode{Content: node.(*ast.Code).Content}
	case *ast.Image:
		n.Value = ImageNode{AltText: node.(*ast.Image).AltText, Url: node.(*ast.Image).URL}
	case *ast.Link:
		n.Value = LinkNode{Text: node.(*ast.Link).Text, Url: node.(*ast.Link).URL}
	case *ast.AutoLink:
		n.Value = AutoLinkNode{Url: node.(*ast.AutoLink).URL, IsRawText: node.(*ast.AutoLink).IsRawText}
	case *ast.Tag:
		n.Value = TagNode{Content: node.(*ast.Tag).Content}
	case *ast.Strikethrough:
		n.Value = StrikethroughNode{Content: node.(*ast.Strikethrough).Content}
	case *ast.EscapingCharacter:
		n.Value = EscapingCharacterNode{Symbol: node.(*ast.EscapingCharacter).Symbol}
	case *ast.Math:
		n.Value = MathNode{Content: node.(*ast.Math).Content}
	case *ast.Highlight:
		n.Value = HighlightNode{Content: node.(*ast.Highlight).Content}
	case *ast.Subscript:
		n.Value = SubscriptNode{Content: node.(*ast.Subscript).Content}
	case *ast.Superscript:
		n.Value = SuperscriptNode{Content: node.(*ast.Superscript).Content}
	case *ast.ReferencedContent:
		n.Value = ReferencedContentNode{ResourceName: node.(*ast.ReferencedContent).ResourceName, Params: node.(*ast.ReferencedContent).Params}
	case *ast.Spoiler:
		n.Value = SpoilerNode{Content: node.(*ast.Spoiler).Content}
	default:
		n.Value = TextNode{Content: ""}
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
	valueBytes, _ := json.Marshal(node.Value)
	switch node.Type {
	case ast.LineBreakNode:
		return &ast.LineBreak{}
	case ast.ParagraphNode:
		paragraphNode := ParagraphNode{}
		json.Unmarshal(valueBytes, &paragraphNode)
		children := convertToASTNodes(paragraphNode.Children)
		return &ast.Paragraph{Children: children}
	case ast.CodeBlockNode:
		codeBlockNode := CodeBlockNode{}
		json.Unmarshal(valueBytes, &codeBlockNode)
		return &ast.CodeBlock{Language: codeBlockNode.Language, Content: codeBlockNode.Content}
	case ast.HeadingNode:
		headingNode := HeadingNode{}
		json.Unmarshal(valueBytes, &headingNode)
		children := convertToASTNodes(headingNode.Children)
		return &ast.Heading{Level: headingNode.Level, Children: children}
	case ast.HorizontalRuleNode:
		horizontalRuleNode := HorizontalRuleNode{}
		json.Unmarshal(valueBytes, &horizontalRuleNode)
		return &ast.HorizontalRule{Symbol: horizontalRuleNode.Symbol}
	case ast.BlockquoteNode:
		blockquoteNode := BlockquoteNode{}
		json.Unmarshal(valueBytes, &blockquoteNode)
		children := convertToASTNodes(blockquoteNode.Children)
		return &ast.Blockquote{Children: children}
	case ast.OrderedListNode:
		orderedListNode := OrderedListNode{}
		json.Unmarshal(valueBytes, &orderedListNode)
		children := convertToASTNodes(orderedListNode.Children)
		return &ast.OrderedList{Number: orderedListNode.Number, Indent: orderedListNode.Indent, Children: children}
	case ast.UnorderedListNode:
		unorderedListNode := UnorderedListNode{}
		json.Unmarshal(valueBytes, &unorderedListNode)
		children := convertToASTNodes(unorderedListNode.Children)
		return &ast.UnorderedList{Symbol: unorderedListNode.Symbol, Indent: unorderedListNode.Indent, Children: children}
	case ast.TaskListNode:
		taskListNode := TaskListNode{}
		json.Unmarshal(valueBytes, &taskListNode)
		children := convertToASTNodes(taskListNode.Children)
		return &ast.TaskList{Symbol: taskListNode.Symbol, Indent: taskListNode.Indent, Complete: taskListNode.Complete, Children: children}
	case ast.MathBlockNode:
		mathBlockNode := MathBlockNode{}
		json.Unmarshal(valueBytes, &mathBlockNode)
		return &ast.MathBlock{Content: mathBlockNode.Content}
	case ast.TableNode:
		tableNode := TableNode{}
		json.Unmarshal(valueBytes, &tableNode)
		rows := [][]string{}
		for _, row := range tableNode.Rows {
			rows = append(rows, row.Cells)
		}
		return &ast.Table{Header: tableNode.Header, Delimiter: tableNode.Delimiter, Rows: rows}
	case ast.EmbeddedContentNode:
		embeddedContentNode := EmbeddedContentNode{}
		json.Unmarshal(valueBytes, &embeddedContentNode)
		return &ast.EmbeddedContent{ResourceName: embeddedContentNode.ResourceName, Params: embeddedContentNode.Params}
	case ast.TextNode:
		textNode := TextNode{}
		json.Unmarshal(valueBytes, &textNode)
		return &ast.Text{Content: textNode.Content}
	case ast.BoldNode:
		boldNode := BoldNode{}
		json.Unmarshal(valueBytes, &boldNode)
		children := convertToASTNodes(boldNode.Children)
		return &ast.Bold{Symbol: boldNode.Symbol, Children: children}
	case ast.ItalicNode:
		italicNode := ItalicNode{}
		json.Unmarshal(valueBytes, &italicNode)
		return &ast.Italic{Symbol: italicNode.Symbol, Content: italicNode.Content}
	case ast.BoldItalicNode:
		boldItalicNode := BoldItalicNode{}
		json.Unmarshal(valueBytes, &boldItalicNode)
		return &ast.BoldItalic{Symbol: boldItalicNode.Symbol, Content: boldItalicNode.Content}
	case ast.CodeNode:
		codeNode := CodeNode{}
		json.Unmarshal(valueBytes, &codeNode)
		return &ast.Code{Content: codeNode.Content}
	case ast.ImageNode:
		imageNode := ImageNode{}
		json.Unmarshal(valueBytes, &imageNode)
		return &ast.Image{AltText: imageNode.AltText, URL: imageNode.Url}
	case ast.LinkNode:
		linkNode := LinkNode{}
		json.Unmarshal(valueBytes, &linkNode)
		return &ast.Link{Text: linkNode.Text, URL: linkNode.Url}
	case ast.AutoLinkNode:
		autoLinkNode := AutoLinkNode{}
		json.Unmarshal(valueBytes, &autoLinkNode)
		return &ast.AutoLink{URL: autoLinkNode.Url, IsRawText: autoLinkNode.IsRawText}
	case ast.TagNode:
		tagNode := TagNode{}
		json.Unmarshal(valueBytes, &tagNode)
		return &ast.Tag{Content: tagNode.Content}
	case ast.StrikethroughNode:
		strikethroughNode := StrikethroughNode{}
		json.Unmarshal(valueBytes, &strikethroughNode)
		return &ast.Strikethrough{Content: strikethroughNode.Content}
	case ast.EscapingCharacterNode:
		escapingCharacterNode := EscapingCharacterNode{}
		json.Unmarshal(valueBytes, &escapingCharacterNode)
		return &ast.EscapingCharacter{Symbol: escapingCharacterNode.Symbol}
	case ast.MathNode:
		mathNode := MathNode{}
		json.Unmarshal(valueBytes, &mathNode)
		return &ast.Math{Content: mathNode.Content}
	case ast.HighlightNode:
		highlightNode := HighlightNode{}
		json.Unmarshal(valueBytes, &highlightNode)
		return &ast.Highlight{Content: highlightNode.Content}
	case ast.SubscriptNode:
		subscriptNode := SubscriptNode{}
		json.Unmarshal(valueBytes, &subscriptNode)
		return &ast.Subscript{Content: subscriptNode.Content}
	case ast.SuperscriptNode:
		superscriptNode := SuperscriptNode{}
		json.Unmarshal(valueBytes, &superscriptNode)
		return &ast.Superscript{Content: superscriptNode.Content}
	case ast.ReferencedContentNode:
		referencedContentNode := ReferencedContentNode{}
		json.Unmarshal(valueBytes, &referencedContentNode)
		return &ast.ReferencedContent{ResourceName: referencedContentNode.ResourceName, Params: referencedContentNode.Params}
	case ast.SpoilerNode:
		spoilerNode := SpoilerNode{}
		json.Unmarshal(valueBytes, &spoilerNode)
		return &ast.Spoiler{Content: spoilerNode.Content}
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
