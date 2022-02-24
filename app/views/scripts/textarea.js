// textareaの高さを変える
let inputId = document.getElementById("input");
let outputId = document.getElementById("output");
let defaultHeight = inputId.clientHeight;

inputId.addEventListener("input", ()=>{
      adjustHeight(inputId)
      outputId.style.height = inputId.style.height;
})

function adjustHeight(element) {
      element.style.height = defaultHeight + "px";
      let sh = element.scrollHeight;
      element.style.height = sh + "px";
}