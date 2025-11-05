<script>
    import Layout from "../components/Layout.svelte";
    import { link } from "svelte5-router";
    import PostCard from "../components/PostCard.svelte";
    import { getRecentPosts } from "../data/posts.js";

    // Get 3 most recent posts
    const recentPosts = getRecentPosts(3);
</script>

<Layout>
    <section class="hero">
        <div class="container">
            <div class="hero-content">
                <h1>Witaj na moim blogu!</h1>
                <p>
                    Prowadzę kanał YouTube, gdzie dzielę się swoją wiedzą na
                    temat programowania w Go, administracji serwerami i
                    budowania infrastruktury IT. Tutaj znajdziesz szczegółowe
                    artykuły i materiały uzupełniające moje filmy.
                </p>
                <div class="hero-actions">
                    <a use:link href="/posty" class="btn-primary"
                        >Zobacz wszystkie posty</a
                    >
                    <a
                        href="https://www.youtube.com/@PragmatycznyPiotr"
                        target="_blank"
                        rel="noopener noreferrer"
                        class="btn-secondary">Obejrzyj na YouTube</a
                    >
                </div>
            </div>
        </div>
    </section>

    <section class="recent-posts">
        <div class="container">
            <h2>Najnowsze posty</h2>
            <div class="posts-grid">
                {#each recentPosts as post (post.id)}
                    <PostCard {post} />
                {/each}
            </div>
            <div class="view-all">
                <a use:link href="/posty" class="view-all-btn"
                    >Zobacz wszystkie posty</a
                >
            </div>
        </div>
    </section>

    <section class="about-preview">
        <div class="container">
            <div class="about-content">
                <div class="about-text">
                    <h2>O mnie</h2>
                    <p>
                        Na moim kanale YouTube dokumentuję swoją naukę języka Go
                        i dzielę się praktycznymi doświadczeniami z codziennego
                        programowania. Pokazuję nie tylko gotowe rozwiązania,
                        ale również proces odkrywania, testowania i wdrażania
                        nowych technik oraz narzędzi. Wszystko, czego się
                        nauczę, staram się przekazywać dalej — prosto i
                        konkretne.
                    </p>
                    <p>
                        Moje główne obszary zainteresowań obejmują cały proces
                        budowania oraz administrowania aplikacjami:
                    </p>
                    <ul>
                        <li>Programowanie (PHP, Go, python, js)</li>
                        <li>Projektowanie architektury aplikacji webowych</li>
                        <li>Implementacja backendu i frontendu</li>
                        <li>Konteneryzacja i automatyzacja (Docker, CI/CD)</li>
                        <li>Konfiguracja oraz utrzymanie infrastruktury serwerowej na Linux</li>
                    </ul>
                    <p><strong>Jeśli chcesz zobaczyć, jak przebiega rzeczywisty proces — od pomysłu, przez kod, po działającą usługę na serwerze — to jesteś w dobrym miejscu!</strong>
                    </p>
                </div>
            </div>
        </div>
    </section>
</Layout>

<style lang="scss">
    .hero {
        background: linear-gradient(
            135deg,
            var(--clr-primary) 0%,
            var(--clr-secondary) 100%
        );
        color: white;
        padding: 100px 0;
        text-align: center;

        &-content {
            max-width: 800px;
            margin: 0 auto;
        }

        h1 {
            font-size: 3.5rem;
            margin-bottom: 20px;
            font-weight: 700;
        }

        p {
            font-size: 1.3rem;
            margin-bottom: 40px;
            opacity: 0.9;
            line-height: 1.6;
        }
    }

    .hero-actions {
        display: flex;
        gap: 20px;
        justify-content: center;
        flex-wrap: wrap;

        .btn-primary {
            background: white;
            color: var(--clr-primary);
            padding: 15px 30px;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            font-size: 1.1rem;
            transition:
                transform 0.2s ease,
                box-shadow 0.2s ease;

            &:hover {
                transform: translateY(-2px);
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
            }
        }

        .btn-secondary {
            background: transparent;
            color: white;
            padding: 15px 30px;
            border: 2px solid white;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            font-size: 1.1rem;
            transition:
                background-color 0.2s ease,
                transform 0.2s ease;

            &:hover {
                background: rgba(255, 255, 255, 0.1);
                transform: translateY(-2px);
            }
        }
    }

    .recent-posts {
        padding: 80px 0;
        background: var(--clr-background);

        h2 {
            text-align: center;
            font-size: 2.5rem;
            color: var(--clr-primary);
            margin-bottom: 50px;
        }

        .posts-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
            gap: 30px;
            max-width: 1200px;
            margin: 0 auto;
        }
    }

    .view-all {
        text-align: center;
        margin-top: 50px;

        .view-all-btn {
            display: inline-block;
            background: var(--clr-primary);
            color: white;
            padding: 15px 30px;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            font-size: 1.1rem;
            transition:
                background-color 0.2s ease,
                transform 0.2s ease;

            &:hover {
                background: var(--clr-secondary);
                transform: translateY(-2px);
            }
        }
    }

    .about-preview {
        padding: 80px 0;
        text-align: center;

        h2 {
            font-size: 2.5rem;
            color: var(--clr-primary);
            margin-bottom: 30px;
        }

        p {
            font-size: 1.1rem;
            line-height: 1.7;
            color: #333;
            max-width: 800px;
            margin: 0 auto 20px;
        }

        ul {
            text-align: left;
            max-width: 600px;
            margin: 0 auto 30px;
            font-size: 1.1rem;
            line-height: 1.7;

            li {
                margin-bottom: 10px;
                color: #333;
            }
        }

        .about-link {
            display: inline-block;
            color: var(--clr-primary);
            text-decoration: none;
            font-weight: 600;
            font-size: 1.1rem;
            padding-bottom: 5px;
            border-bottom: 2px solid var(--clr-primary);
            transition:
                color 0.2s ease,
                border-color 0.2s ease;

            &:hover {
                color: var(--clr-secondary);
                border-bottom-color: var(--clr-secondary);
            }
        }
    }

    @media (max-width: 768px) {
        .hero {
            padding: 60px 0;

            h1 {
                font-size: 2.5rem;
            }

            p {
                font-size: 1.1rem;
            }
        }

        .hero-actions {
            flex-direction: column;
            align-items: center;

            .btn-primary,
            .btn-secondary {
                min-width: 250px;
                text-align: center;
            }
        }

        .recent-posts {
            padding: 60px 0;

            h2 {
                font-size: 2rem;
            }
        }

        .posts-grid {
            grid-template-columns: 1fr;
            gap: 20px;
        }

        .about-preview {
            padding: 60px 0;

            h2 {
                font-size: 2rem;
            }
        }
    }
</style>
