<script>
    import Layout from "../components/Layout.svelte";
    import { link } from "svelte5-router";
    import { getPostBySlug, formatDate } from "../data/posts.js";

	let { slug } = $props();

    let post = getPostBySlug(slug);
</script>

<Layout>
    {#if post}
        <article class="post-container">
            <header class="post-header">
                <div class="post-meta">
                    <time datetime={post.publishDate}>{formatDate(post.publishDate)}</time>
                </div>
                <h1 class="post-title">{post.title}</h1>
                <div class="post-tags">
                    {#each post.tags as tag}
                        <span class="tag">{tag}</span>
                    {/each}
                </div>
            </header>

            <div class="post-content">
                {@html post.content}
            </div>

            <footer class="post-footer">
                <nav class="post-navigation">
                    <a use:link href="/posty" class="back-to-posts">
                        ← Wszystkie posty
                    </a>
                </nav>
            </footer>
        </article>
    {:else}
        <div class="container">
            <section class="not-found">
                <h1>Post nie znaleziony</h1>
                <p>Niestety, post o podanym ID nie istnieje.</p>
                <a use:link href="/posty" class="back-link">Wróć do listy postów</a>
            </section>
        </div>
    {/if}
</Layout>

<style lang="scss">
    .post-container {
        max-width: 800px;
        margin: 0 auto;
        padding: 40px 20px;
        background: white;
        border-radius: 12px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        margin-top: 40px;
        margin-bottom: 40px;
    }

    .post-header {
        margin-bottom: 40px;
        padding-bottom: 30px;
        border-bottom: 2px solid var(--clr-background);
    }

    .post-meta {
        display: flex;
        align-items: center;
        gap: 15px;
        margin-bottom: 20px;
        font-size: 0.9rem;
        color: #666;

        time {
            font-weight: 500;
        }

        .read-time {
            color: var(--clr-secondary);
            font-weight: 500;
        }
    }

    .post-title {
        font-size: 2.5rem;
        line-height: 1.2;
        color: var(--clr-primary);
        margin: 0 0 20px 0;
        font-weight: 700;
    }

    .post-tags {
        display: flex;
        flex-wrap: wrap;
        gap: 8px;

        .tag {
            background: var(--clr-background);
            color: var(--clr-primary);
            padding: 6px 16px;
            border-radius: 20px;
            font-size: 0.85rem;
            font-weight: 500;
        }
    }

    .post-content {
        line-height: 1.7;
        font-size: 1.1rem;
        color: #333;

        h1, h2, h3, h4 {
            color: var(--clr-primary);
            margin-top: 40px;
            margin-bottom: 20px;
        }

        h2 {
            font-size: 1.8rem;
            border-bottom: 2px solid var(--clr-background);
            padding-bottom: 10px;
        }

        h3 {
            font-size: 1.4rem;
        }

        p {
            margin-bottom: 20px;
        }

        ul, ol {
            margin: 20px 0;
            padding-left: 30px;

            li {
                margin-bottom: 10px;
            }
        }

        code {
            background: #f4f4f4;
            padding: 2px 6px;
            border-radius: 4px;
            font-family: 'Courier New', monospace;
            font-size: 0.9rem;
        }

        pre {
            background: #2d2d2d;
            color: #f8f8f2;
            padding: 20px;
            border-radius: 8px;
            overflow-x: auto;
            margin: 30px 0;

            code {
                background: none;
                padding: 0;
                font-size: 1rem;
            }
        }

        blockquote {
            border-left: 4px solid var(--clr-primary);
            padding-left: 20px;
            margin: 30px 0;
            font-style: italic;
            color: #666;
        }

        a {
            color: var(--clr-secondary);
            text-decoration: none;
            border-bottom: 1px solid transparent;
            transition: border-color 0.2s ease;

            &:hover {
                border-bottom-color: var(--clr-secondary);
            }
        }

        /* Responsive YouTube Embed */
        :global(.video-container) {
            position: relative;
            padding-bottom: 56.25%; /* 16:9 Aspect Ratio */
            height: 0;
            overflow: hidden;
            max-width: 100%;
            margin: 30px 0;
        }

        :global(.video-container iframe) {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            border: none;
        }

        /* Direct iframe styling for YouTube */
        :global(iframe[src*="youtube.com"]) {
            width: 100%;
            height: auto;
            aspect-ratio: 16/9;
            border: none;
            margin: 30px 0;
            border-radius: 8px;
        }
    }

    .post-footer {
        margin-top: 60px;
        padding-top: 30px;
        border-top: 2px solid var(--clr-background);
    }

    .video-section {
        background: linear-gradient(135deg, #ff6b6b 0%, #ee5a24 100%);
        color: white;
        padding: 30px;
        border-radius: 12px;
        margin-bottom: 40px;
        text-align: center;

        h3 {
            margin-top: 0;
            color: white;
            font-size: 1.5rem;
        }

        p {
            margin-bottom: 20px;
            opacity: 0.9;
        }

        .video-button {
            display: inline-block;
            background: white;
            color: #ff6b6b;
            padding: 12px 24px;
            border-radius: 6px;
            text-decoration: none;
            font-weight: 600;
            transition: transform 0.2s ease, box-shadow 0.2s ease;

            &:hover {
                transform: translateY(-2px);
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
            }
        }
    }

    .post-navigation {
        text-align: center;

        .back-to-posts {
            color: var(--clr-primary);
            text-decoration: none;
            font-weight: 600;
            font-size: 1.1rem;
            transition: color 0.2s ease;

            &:hover {
                color: var(--clr-secondary);
            }
        }
    }

    .not-found {
        text-align: center;
        padding: 80px 20px;

        h1 {
            font-size: 2.5rem;
            color: var(--clr-primary);
            margin-bottom: 20px;
        }

        p {
            font-size: 1.2rem;
            color: #666;
            margin-bottom: 30px;
        }

        .back-link {
            display: inline-block;
            background: var(--clr-primary);
            color: white;
            padding: 12px 24px;
            border-radius: 6px;
            text-decoration: none;
            font-weight: 600;
            transition: background-color 0.2s ease;

            &:hover {
                background: var(--clr-secondary);
            }
        }
    }

    @media (max-width: 768px) {
        .post-container {
            margin: 20px;
            padding: 30px 20px;
        }

        .post-title {
            font-size: 2rem;
        }

        .post-content {
            font-size: 1rem;
        }

        .video-section {
            padding: 20px;
        }
    }
</style>