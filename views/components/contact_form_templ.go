// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.639
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ContactForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"grid gap-2\"><div class=\"grid grid-cols-2 gap-2\"><label for=\"first_name\" class=\"sr-only\">First Name</label> <input type=\"text\" name=\"first_name\" id=\"first_name\" class=\"contact-input\" placeholder=\"* First Name\" required> <label for=\"last_name\" class=\"sr-only\">Last Name</label> <input type=\"text\" name=\"last_name\" id=\"last_name\" class=\"contact-input\" placeholder=\"* Last Name\" required></div><label for=\"email\" class=\"sr-only\">Email</label> <input type=\"email\" name=\"email\" id=\"email\" class=\"contact-input\" placeholder=\"* Email\" required> <textarea name=\"message\" id=\"message\" class=\"resize-none contact-input\" rows=\"5\" placeholder=\"* Message\" required></textarea> <input type=\"submit\" value=\"Send Enquiry\" class=\"bg-orange text-white contact-input\"><style>\n            .contact-input {\n                padding-inline: 0.5rem;\n                padding-block: 0.25rem;\n                border-radius: 1rem;\n            }\n        </style></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
