let inputId = document.getElementById("input");
let defaultHeight = inputId.clientHeight;

document.addEventListener("input", ()=>{
      inputId.style.height = defaultHeight + "px";
      let sh = inputId.scrollHeight;
      inputId.style.height = sh + "px";
})
