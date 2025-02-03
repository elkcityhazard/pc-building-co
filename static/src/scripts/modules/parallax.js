export const Parallax = function() {
    const bgImage = document.getElementById('homeHeader');

    if (!bgImage) {
        return; 
    }

    const updateParallax = () => {
        const totalHeight = document.body.scrollHeight;
        const currentScroll = document.documentElement.scrollTop;

        if (currentScroll > bgImage.getBoundingClientRect().height) {
            return;
        }

        let p = 50 - (currentScroll / totalHeight * 200);
        bgImage.style.backgroundPositionY = p + "%";
    };
    
    const observer = new IntersectionObserver(entries => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                window.addEventListener('scroll', updateParallax);
                updateParallax(); // Initial call to set the position
            } else {
                window.removeEventListener('scroll', updateParallax);
            }
        });
    });

    observer.observe(bgImage);
};
