document.addEventListener("DOMContentLoaded", function () {
    const openButton = document.getElementById("sidebar-toggle");
    const closeButton = document.getElementById("sidebar-close");
    const sidebar = document.getElementById("sidebar");

    openButton.addEventListener("click", function () {
        sidebar.classList.add("open");
    });

    closeButton.addEventListener("click", function () {
        sidebar.classList.remove("open");
    });
});

