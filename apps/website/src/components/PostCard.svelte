<script>
    import { link } from "svelte5-router";
    import { formatDate } from "../data/posts.js";

    let { post } = $props();
</script>

<article class="post-card">
    {#if post.thumbnail}
        <div class="post-thumbnail">
            <img src={post.thumbnail} alt={post.title} loading="lazy"/>
        </div>
    {/if}
    <div class="post-content">
        <header class="post-header">
            <div class="post-meta">
                <time datetime={post.publishDate}>{formatDate(post.publishDate)}</time>
            </div>
            <h2 class="post-title">
                <a use:link href={`/post/${post.slug}`}>{post.title}</a>
            </h2>
        </header>

        <p class="post-excerpt">{post.excerpt}</p>

        <div class="post-tags">
            {#each post.tags as tag}
                <span class="tag">{tag}</span>
            {/each}
        </div>

        <div class="post-actions">
            <a use:link href={`/post/${post.slug}`} class="read-more">
                Czytaj dalej →
            </a>
        </div>
    </div>
</article>

<style lang="scss">
    .post-card {
        background: white;
        border-radius: 12px;
        overflow: hidden;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        transition: box-shadow 0.2s ease;
        display: flex;
        flex-direction: column;

        &:hover {
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
        }
    }

    .post-thumbnail {
        width: 100%;
        overflow: hidden;

        img {
            width: 100%;
            height: auto;
            object-fit: contain;
        }
    }

    .post-content {
        padding: 30px;
        flex: 1;
        display: flex;
        flex-direction: column;
    }

    .post-header {
        margin-bottom: 20px;
    }

    .post-meta {
        display: flex;
        align-items: center;
        gap: 15px;
        margin-bottom: 10px;
        font-size: 0.9rem;
        color: #666;

        time {
            font-weight: 500;
        }
    }

    .post-title {
        margin: 0;
        font-size: 1.4rem;
        line-height: 1.3;

        a {
            color: var(--clr-primary);
            text-decoration: none;
            transition: color 0.2s ease;

            &:hover {
                color: var(--clr-secondary);
            }
        }
    }

    .post-excerpt {
        font-size: 1rem;
        line-height: 1.6;
        color: #333;
        margin-bottom: 20px;
    }

    .post-tags {
        display: flex;
        flex-wrap: wrap;
        gap: 8px;
        margin-bottom: 20px;

        .tag {
            background: var(--clr-background);
            color: var(--clr-primary);
            padding: 4px 12px;
            border-radius: 20px;
            font-size: 0.85rem;
            font-weight: 500;
        }
    }

    .post-actions {
        margin-top: auto;

        .read-more {
            color: var(--clr-primary);
            text-decoration: none;
            font-weight: 600;
            font-size: 0.95rem;
            transition: color 0.2s ease;

            &:hover {
                color: var(--clr-secondary);
            }
        }
    }

    /* Responsive adjustments */
    @media (max-width: 768px) {
        .post-title {
            font-size: 1.2rem;
        }

        .post-excerpt {
            font-size: 0.95rem;
        }

        .post-content {
            padding: 20px;
        }
    }
</style>