function formatTextarea(textareaID = "") {
  this.textareaID = textareaID;
  this.textarea = document.getElementById(this.textareaID) || null;
  this.initialValue = this.textarea.value;
  this.currentValue = "";
  if (!this.textarea) return null;

  // clear the box and reinstate message if needed
  this.clearBox = function () {
    this.textarea.addEventListener(
      "focus",
      function () {
        if (this.currentValue != "") {
          this.textarea.value = this.currentValue;
          return;
        }
        this.currentvalue = "";
        this.textarea.value = this.currentvalue;
      }.bind(this),
    );

    // handle what to do on focus out (keep text, reset default msg)
    this.textarea.addEventListener(
      "focusout",
      function () {
        if (this.textarea.value != "") return;
        this.textarea.value =
          this.currentValue != "" ? this.currentValue : this.initialValue;
      }.bind(this),
    );
    // handle textarea updates
    this.textarea.addEventListener(
      "change",
      function (e) {
        this.currentValue = e.target.value;
      }.bind(this),
    );
  };

  this.Init = function () {
    // make sure dom content is loaded
    document.addEventListener("DOMContentLoaded", this.clearBox.bind(this));
  };

  // called on  new clearTextarea(id)
  return this.Init();
}

export { formatTextarea };
