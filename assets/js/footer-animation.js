const animatedElements = document.querySelectorAll('.animate-in');

if (animatedElements.length) {
  const observer = new IntersectionObserver(
    (entries, obs) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible');
          obs.unobserve(entry.target); // animate once
        }
      });
    },
    {
      threshold: 0.2
    }
  );

  animatedElements.forEach(el => observer.observe(el));
}

const observer = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      entry.target.classList.add("visible");
      observer.unobserve(entry.target); // animate only once
    }
  });
}, { threshold: 0.2 });

document.querySelectorAll(".fade-in").forEach(el => observer.observe(el));

const footer = document.querySelector("footer");
const footerParallaxEls = footer.querySelectorAll("[data-footer-speed]");
let footerTicking = false;

function clamp(v, min = 0, max = 1) {
  return Math.min(max, Math.max(min, v));
}

function onFooterScroll() {
  const scrollY = window.scrollY;
  const vh = window.innerHeight;

  const footerTop = footer.offsetTop;
  const footerHeight = footer.offsetHeight;

  const maxScroll =
    document.documentElement.scrollHeight - vh;

  // progress starts when footer enters viewport
  // ends when page hits max scroll
  const start = footerTop - vh;
  const end = maxScroll;

  const progress = clamp(
    (scrollY - start) / (end - start)
  );

  footerParallaxEls.forEach(el => {
    const speed = parseFloat(el.dataset.footerSpeed);
    const maxOffset = footerHeight * speed;

    // interpolate from -maxOffset → 0
    const translateY = (1 - progress) * -maxOffset;

    el.style.transform = `translateY(${translateY}px)`;
  });

  footerTicking = false;
}

window.addEventListener("scroll", () => {
  if (!footerTicking) {
    requestAnimationFrame(onFooterScroll);
    footerTicking = true;
  }
});

window.addEventListener("resize", onFooterScroll);
onFooterScroll();