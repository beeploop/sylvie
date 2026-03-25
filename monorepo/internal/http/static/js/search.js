(() => {
    const searchForm = document.getElementById("search-form");

    searchForm.addEventListener("submit", (e) => {
        e.preventDefault();

        /** @type {string} input */
        const input = searchForm.querySelector("#input").value;
        if (input.trim() === "") return;

        searchForm.submit();
    })
})();
