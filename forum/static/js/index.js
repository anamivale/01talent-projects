function Open(id) {
  let comments = document.getElementById(id);

  comments.classList.toggle("close");
}
function likePost(postID) {
  document.alert("welcome");
  fetch(`/like`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ postID: postID }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        let x = "likes " + data.likeCount;
        document.getElementById(`like-count-${postID}`).textContent = x;
      } else {
        alert(data.message);
      }
    })
    .catch((error) => console.error("Error:", error));
}

function disLikePost(postID, userID) {
  fetch(`/dislike`, {
    method: "POST",
  });
}
