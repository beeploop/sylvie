(() => {
    const video = document.getElementById("video_player");
    const videoURL = video.getAttribute("data-source");
    const qualitySelector = document.getElementById("quality-selector");

    if (Hls.isSupported()) {
        console.log("hls is supported");

        const hls = new Hls();
        hls.loadSource(videoURL);
        hls.attachMedia(video);

        hls.on(Hls.Events.MANIFEST_PARSED, () => {
            const levels = [...hls.levels].sort((a, b) => b.height - a.height);
            levels.forEach((level, index) => {
                const option = document.createElement("option");
                option.value = index;
                option.textContent = `${level.height}p`;
                qualitySelector.appendChild(option);
            });
        });

        hls.on(Hls.Events.FRAG_CHANGED, (_, data) => {
            const level = hls.levels[data.frag.level];

            console.log({
                quality: `${level.height}p`,
                bitrate: `${Math.round(level.bitrate / 1000)} kbps`,
                segment: data.frag.relurl,
            });
        });

        hls.on(Hls.Events.LEVEL_CHANGED, (_, data) => {
            qualitySelector.value = String(data.level);
        });

        qualitySelector.addEventListener("change", (e) => {
            const level = parseInt(e.target.value, 10);
            console.log({ level });

            if (level === -1) {
                hls.currentLevel = -1;
                return;
            }

            hls.nextLevel = level;
        });
    } else if (video.canPlayType('application/vnd.apple.mpegurl')) {
        console.log("hls is not supported");

        video.src = videoURL;
    }
})();
