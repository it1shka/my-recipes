// making auto-resizing text-areas
const textareas = document.getElementsByClassName('autoresize-textarea')
for(const tx of textareas) {
  const style = `height:${tx.scrollHeight}px;overflow-y:hidden;`
  tx.setAttribute('style', style)
  tx.addEventListener('input', function() {
    this.style.height = 'auto'
    this.style.height = `${this.scrollHeight}px`
  }) 
}

// making auto-resizing inputs
// const inputs = document.getElementsByClassName('autoresize-input')
// for(const inp of inputs) {
//   inp.style.width = '50px'
//   inp.addEventListener('input', function() {
//     this.style.width = `${this.value.length}ch`
//   })
// }