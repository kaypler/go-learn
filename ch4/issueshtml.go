// go run ./issueshtml.go repo:golang/go commenter:gopherbot json encoder > issues.html
package main

import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}}</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</td>
</tr>
{{end}}
</table>	
`))