
async function layoutPhotoWall() {
    const wall = document.querySelector('.photo-wall');
    if (!wall) return;

    const images = wall.querySelectorAll('img');
    const columns = window.innerWidth >= 768 ? 3 : 2;
    const gap = 16;
    const totalWidth = wall.clientWidth;
    if (totalWidth === 0) return;

    const colWidth = (totalWidth - (gap * (columns - 1))) / columns;
    let colHeights = Array(columns).fill(0);
    let maxBottom = 0;

    // Loop through images and ensure we wait for their real dimensions
    for (const img of images) {
        img.style.width = `${colWidth}px`;
        img.style.position = 'absolute';

        // Wait for the image to load its true natural dimensions
        if (!img.complete || img.naturalWidth === 0) {
            try {
                await img.decode();
            } catch (e) {
                // Fallback if decode fails
            }
        }

        const shortestCol = colHeights.indexOf(Math.min(...colHeights));
        const x = shortestCol * (colWidth + gap);
        const y = colHeights[shortestCol];

        img.style.left = `${x}px`;
        img.style.top = `${y}px`;

        // Calculate exact height using real natural dimensions
        let imgHeight = colWidth * 1.25; // fallback
        if (img.naturalWidth && img.naturalHeight) {
            imgHeight = colWidth * (img.naturalHeight / img.naturalWidth);
        }

        colHeights[shortestCol] += imgHeight + gap;

        const imageBottom = y + imgHeight;
        if (imageBottom > maxBottom) {
            maxBottom = imageBottom;
        }
    }

    // Apply the correct full height so the footer is pushed down properly
    wall.style.height = `${maxBottom + gap}px`;
}

// Run on load, resize, and DOM ready
window.addEventListener('load', layoutPhotoWall);
window.addEventListener('resize', layoutPhotoWall);
document.addEventListener('DOMContentLoaded', layoutPhotoWall);

// Trigger immediately
layoutPhotoWall();
