let btn = document.querySelector(".open-comment-section")

btn.addEventListener("click", (e)=>{
   let comments = document.querySelector(".comment-section")

    comments.classList.toggle("close")
})