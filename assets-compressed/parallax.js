const parallaxEls = document.querySelectorAll('.parallax');
let ticking = false;

function onScroll() {
  const scrollY = window.scrollY;
  parallaxEls.forEach(el => {
    const speed = parseFloat(el.dataset.speed);
    const direction = el.dataset.direction === 'down' ? 1 : -1;
    const limitValue = el.dataset.limit ? parseFloat(el.dataset.limit) : null;
    
    // Calculate limit based on viewport width if limit exists
    let limit = null;
    if (limitValue !== null) {
      limit = (window.innerWidth / 100) * limitValue;
    }
    
    let translateY = scrollY * speed * direction;
    
    // Apply limit if it exists
    if (limit !== null) {
      if (direction === -1) {
        translateY = Math.max(translateY, -limit);
      } else {
        translateY = Math.min(translateY, limit);
      }
    }
    
    el.style.transform = `translateY(${translateY}px)`;
  });
  ticking = false;
}

window.addEventListener('scroll', () => {
  if (!ticking) {
    window.requestAnimationFrame(onScroll);
    ticking = true;
  }
});