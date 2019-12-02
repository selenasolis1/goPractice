//see page 133
//outline prints the outline of an HTML document tree
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

//forEachNode calls the functions pre(x) and post(x) for each node
//x in the tree rooted at n. Both functions are optional
//pre is called before the children are visited (preorder) and
//post is called after (postorder)
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	fmt.Println(n.Data)
	fmt.Println(pre)
	fmt.Println(post)
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {

	if n.Type == html.ElementNode && len(n.Attr) == 0 {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	} else if n.Type == html.ElementNode {
		var str string
		for _, a := range n.Attr {
			str += " " + a.Key + "=" + a.Val
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, str)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
