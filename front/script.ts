async function fetchAlbum() {
    try {
        const response = await fetch("http://localhost:8080/album-of-the-day");
        const album = await response.json();

        const titleElement = document.getElementById("album-title");
        const artistElement =   document.getElementById("album-artist");
        const yearElement = document.getElementById("album-year");
        
        if (titleElement) {
            titleElement.innerText = album.album || "Unknown album";
        }
        if (artistElement) {
            artistElement.innerText = album.artist || "Unknown Artist";
        }
        if (yearElement) {
            yearElement.innerText = "Released: " + album.release_year;
        }
        
        const cover = document.getElementById("album-cover") as HTMLImageElement;
        if (cover) {
            cover.src = album.cover;
            cover.style.display = "block";
        }
    }
    catch (error) {
        console.error("Error fetching album: ", error);
        const album = document.getElementById("album-title")
        if (album) {
            album.innerText = "Could not load album."
        }
    }
}


fetchAlbum();