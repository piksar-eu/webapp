<script>
    import Subpage from "../components/Subpage.svelte";
    import { onMount, onDestroy } from "svelte";
    import 'leaflet/dist/leaflet.css';

    let map;

    onMount(async () => {
        if (typeof window !== 'undefined') {
            const L = await import('leaflet');

            map = L.map('map').setView([50.041187, 21.999121], 13);

            L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
            }).addTo(map);

            L.marker([50.041187, 21.999121])
                .addTo(map)
                .bindPopup('Lokalizacja mojej firmy.')
                .openPopup();
        }
    });

    onDestroy(() => {
        if (map) {
            map.remove();
        }
    });
</script>

<Subpage header="Kontakt">
    <div id="map"></div>
    <h2>Moja firma sp. z.o.o.</h2>
    <p>
        ul. Testowa 12<br/>
        66-666 Miasto<br/>
        NIP: 666-66-66-666<br/>
        email: adres@email.pl<br/>
        tel: 666 666 666<br/>
    </p>
</Subpage>

<style lang="scss">
    #map {
        width: 600px;
        height: 400px;
        float: right;
        padding-left: 8px;
        margin-bottom: 8px;
    }

    p {
        margin-bottom: 16px;
    }
</style>