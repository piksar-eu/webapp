<script>
    import Layout from "../components/Layout.svelte";
    import PostCard from "../components/PostCard.svelte";
    import { posts } from "../data/posts.js";
</script>

<Layout>
    <section class="posts-hero">
        <div class="container">
            <h1>Moje Posty</h1>
            <p>Tutaj znajdziesz wszystkie moje artykuły związane z programowaniem w Go, administracją serwerami i infrastrukturą IT. Każdy post to zapis mojego doświadczenia z kanału YouTube.</p>
        </div>
    </section>

    <section class="posts-list container">
        {#each posts as post (post.id)}
            <PostCard {post} />
        {/each}
    </section>
</Layout>

<style lang="scss">
    .posts-hero {
        background: linear-gradient(135deg, var(--clr-primary) 0%, var(--clr-secondary) 100%);
        color: white;
        padding: 80px 0;
        text-align: center;
        margin-bottom: 60px;

        h1 {
            font-size: 3rem;
            margin-bottom: 20px;
            font-weight: 700;
        }

        p {
            font-size: 1.2rem;
            max-width: 800px;
            margin: 0 auto;
            opacity: 0.9;
        }
    }

    .posts-list {
        margin: 0 auto;
        padding: 0 20px;
        margin-bottom: 64px;
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
        gap: 30px;
        align-items: stretch;
    }

    @media (min-width: 992px) {
        .posts-list {
            grid-template-columns: repeat(2, 1fr);
        }
    }

    @media (min-width: 1200px) {
        .posts-list {
            grid-template-columns: repeat(3, 1fr);
        }
    }

    @media (max-width: 768px) {
        .posts-list {
            grid-template-columns: 1fr;
            padding: 0 15px;
        }
    }

    
    
    @media (max-width: 768px) {
        .posts-hero {
            padding: 60px 0;

            h1 {
                font-size: 2.2rem;
            }

            p {
                font-size: 1.1rem;
            }
        }

      }
</style>