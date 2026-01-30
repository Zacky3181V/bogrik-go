const overlay = document.getElementById("overlay");
const overlayImg = document.getElementById("overlay-img");
document.querySelectorAll(".zoomable").forEach(img => {
  img.addEventListener("click", () => {
    overlayImg.src = img.src;
    overlay.style.display = "grid";
  });
});
overlay.addEventListener("click", () => {
  overlay.style.display = "none";
});