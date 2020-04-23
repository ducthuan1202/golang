(function (doc, setting) {

  console.log(doc.title);

  console.log(JSON.stringify(setting))
})(document, {
    width: {{ .Width}},
    height: {{ .Height}},
    text: "{{ .Text}}",
});
