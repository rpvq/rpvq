document.addEventListener("DOMContentLoaded", () => {
    // 1. Typing effect for the main title
    const title = document.getElementById('typing-text');
    const originalText = title.innerText;
    title.innerText = '';
    let index = 0;

    function type() {
        if (index < originalText.length) {
            title.innerText += originalText.charAt(index);
            index++;
            setTimeout(type, 150);
        }
    }
    type();

    // 2. Scroll Reveal Animation
    const reveals = document.querySelectorAll('.reveal');

    function reveal() {
        reveals.forEach(element => {
            const windowHeight = window.innerHeight;
            const elementTop = element.getBoundingClientRect().top;
            const elementVisible = 150;

            if (elementTop < windowHeight - elementVisible) {
                element.classList.add('active');
            }
        });
    }

    window.addEventListener('scroll', reveal);
    reveal(); // Check on load
});
