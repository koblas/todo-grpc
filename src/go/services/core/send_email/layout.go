package send_email

import (
	"html/template"
)

const baseLAYOUT = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<title>{{ .Title }}</title>
		<meta http-equiv="content-type" content="text/html;charset=UTF-8">
		<style type="text/css">
			body {
				background:#f1f1f1;
				margin:0;
				padding:0
			}
			img {
				border:0 !important;
				border:none !important;
				outline:none !important
			}
			a {
				text-decoration:none !important
			}
			a img {
				border:0 !important;
				border:none !important
			}
			table {
				border:0;
				border:none;
				border-collapse:collapse
			}
			table tr {
				border-spacing:0 !important
			}
			table td {
				border-collapse:collapse !important
			}
			.thin {
				font-family:'HelveticaNeue-Thin','HelveticaNeue-Light','Helvetica-Light','Helvetica','Arial','Lucida Grande',sans-serif
			}
			.light {
				font-family:'HelveticaNeue-Light','HelveticaNeue-Thin','Helvetica-Light','Helvetica','Arial','Lucida Grande',sans-serif
			}
			.main-table {
				border:1px solid #e3e3e3;
				max-width:600px !important;
				min-width:320px
			}
			#logo img {
				border:0;
				border:none
			}
			.title {
				color:#333333;
				font-size:34px;
				line-height:120%;
				width:100%
			}
			.text, .text font {
				color:#333333;
				font-size:15px;
				font-weight:normal;
				line-height:160%
			}
			.footer {
				color:#888888;
			}
			.footer, .footer font {
				font-size:11px;
				line-height:170%;
				text-align:center
			}
			.text a, .footer a {
				color:#0097ee !important;
				font-weight:bold !important;
				text-decoration:none !important
			}

			/* SCREEN SIZE ADJUSTMENTS */

			@media only screen and (max-width:600px){
				table {width:100% !important}
 				img.full-width {width:100% !important}
 				.main-table {border-left-width:0 !important;border-right-width:0 !important}
 			}
 			@media only screen and (max-width:550px){
 				.title {font-size:32px !important}
 			}
			@media only screen and (max-width:500px){
				.title {font-size:30px !important}
			}
			@media only screen and (max-width:450px){
				.title {font-size:28px !important}
			}
			@media only screen and (max-width:400px){
				.title {font-size:26px !important}
			}
			@media only screen and (max-width:350px){
				.title {font-size:24px !important}
			}
			@media only screen and (max-width:300px){
				.title {font-size:22px !important}
			}
		</style>
	</head>
	<body>
		<table class="wrapper" bgcolor="#f1f1f1" cellspacing="0" cellpadding="0" align="center" width="100%">
			<tr>
				<td align="center" height="100" valign="middle">
					<a id="logo" href="{{ .LogoUrl }}">

						<!-- img src="LOGO_URL" style="border:0;border:none" width="108" alt="{{ .AppName }}" title="{{ .AppName }}" -->
                        <h2>{{ .AppName }}</h2>
					</a>
				</td>
			</tr>
			<tr>
				<td align="center">
					<table class="main-table" bgcolor="#ffffff" cellspacing="0" cellpadding="0" align="center" width="600">
						<tr>
							<td width="8%">
								<!-- img style="max-width:46px !important" src="" width="100%" height="1" alt="" -->
							</td>
							<td width="84%" align="center">
                                {{ template "content" . }}
							</td>
							<td width="8%">
                               <!-- img style="max-width:46px !important" src="" width="100%" height="1" alt="" -->
                                <a></a>
							</td>
						</tr>
						
					</table>

				</td>
			</tr>
			<tr>
				<td class="footer light" align="center" valign="top">
					<br>
					<font face="'HelveticaNeue-Light','HelveticaNeue-Thin','Helvetica-Light','Helvetica','Arial','Lucida Grande',sans-serif" size="2">
						This email has been sent to you by <a style="color:#404040 !important;font-weight:normal !important;text-decoration:none !important" href="{{ .LogoUrl }}" target="_blank"><font color="#404040">{{ .AppName }}</font></a>.<br>__ADDRESS_XXX__<br><a style="color:#0097ee !important;font-weight:bold !important;text-decoration:none !important" href="{{ .PrivacyUrl }}"><font color="#0097ee">Privacy Policy</font></a><br><br><br><br>
					</font>
				</td>
			</tr>
		</table>
	<img src="{{ .OpenPingURL }}" height="1" width="1"></body>
</html>
`

var _layout = template.Must(template.New("layout").Parse(baseLAYOUT))

// Layout -- exported for testing
func Layout() *template.Template {
	return template.Must(_layout.Clone())
}
