all: invite_user.html password_change.html password_recovery.html register_user.html

%.html: %.mjml _footer.mjml _headblock.mjml _header.mjml
	npx mjml -o $@ $<

