(() => {
    const triggers = document.querySelectorAll("[data-dialog-trigger]");
    const dialogs = document.querySelectorAll("[data-dialog-target]");

    /** @type {Map<string, Element>} targets */
    const dialogMap = new Map();

    for (const dialog of dialogs) {
        const dialogID = dialog.getAttribute("id");
        if (dialogID === "") continue;
        dialogMap.set(dialogID, dialog)
    }

    triggers.forEach((trigger) => {
        trigger.addEventListener("click", () => {
            const triggerID = trigger.getAttribute("data-dialog-trigger");
            dialogMap.get(triggerID)?.showModal();

            const url = new URL(window.location.href);
            url.searchParams.set("dialog", triggerID)
            window.history.replaceState(null, "", url.toString());
        });
    });

    dialogs.forEach((dialog) => {
        const close = dialog.querySelector("#close")
        close?.addEventListener("click", () => {
            dialog.close();

            const url = new URL(window.location.href);
            url.searchParams.delete("dialog")
            window.history.replaceState(null, "", url.toString());
        });
    })

    const queryParams = new URLSearchParams(window.location.search);
    if (queryParams.has("dialog")) {
        const dialogID = queryParams.get("dialog");
        dialogMap.get(dialogID)?.showModal();
    }
})();
