document.addEventListener("DOMContentLoaded", () => {
    const buttons = document.querySelectorAll(".info-btn");

    buttons.forEach(btn => {
        btn.addEventListener("click", () => {
            const content = btn.nextElementSibling;

            document.querySelectorAll(".info-content").forEach(c => {
                if (c !== content) c.classList.remove("active");
            });

            content.classList.toggle("active");
        });
    });
});
