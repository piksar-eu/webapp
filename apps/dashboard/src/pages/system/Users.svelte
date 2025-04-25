<script>
	import Breadcrumbs from "../../components/Breadcrumbs.svelte";
	import Layout from "../../components/Layout.svelte";
	import { getUsers, getRoles } from "../../api";
    import { link } from "svelte5-router";
    import { onMount } from "svelte";

	let roles = $state([])
	let users = $state([])
	onMount(async () => {
		[roles, users] = await Promise.all([
			getRoles(),
			getUsers()
		]);
	})

	const roleName = (id) => {
		return roles.find(r => r.id === id)?.name
	}
</script>

<Layout>
	{#snippet header()}
		<Breadcrumbs items={[{href: '/system', name: 'System'}, {name: 'Użytkownicy'}]}></Breadcrumbs>
	{/snippet}
	{#snippet main()}
		<div class="content grid">
			<svg xmlns="http://www.w3.org/2000/svg" style="display: none;">
				<!-- Font Awesome Edit Icon -->
				<symbol id="fa-edit-icon" viewBox="0 0 512 512">
					<path d="M441 58.9L453.1 71c9.4 9.4 9.4 24.6 0 33.9L424 134.1 377.9 88 407 58.9c9.4-9.4 24.6-9.4 33.9 0zM209.8 256.2L344 121.9 390.1 168 255.8 302.2c-2.9 2.9-6.5 5-10.4 6.1l-58.5 16.7 16.7-58.5c1.1-3.9 3.2-7.5 6.1-10.4zM373.1 25L175.8 222.2c-8.7 8.7-15 19.4-18.3 31.1l-28.6 100c-2.4 8.4-.1 17.4 6.1 23.6s15.2 8.5 23.6 6.1l100-28.6c11.8-3.4 22.5-9.7 31.1-18.3L487 138.9c28.1-28.1 28.1-73.7 0-101.8L474.9 25C446.8-3.1 401.2-3.1 373.1 25zM88 64C39.4 64 0 103.4 0 152L0 424c0 48.6 39.4 88 88 88l272 0c48.6 0 88-39.4 88-88l0-112c0-13.3-10.7-24-24-24s-24 10.7-24 24l0 112c0 22.1-17.9 40-40 40L88 464c-22.1 0-40-17.9-40-40l0-272c0-22.1 17.9-40 40-40l112 0c13.3 0 24-10.7 24-24s-10.7-24-24-24L88 64z"/>
				</symbol>
			</svg>
			<table>
				<thead>
				<tr>
					<th scope="col">Email</th>
					<th scope="col">Nazwa</th>
					<th scope="col">Role</th>
					<th scope="col"></th>
				</tr>
				</thead>
				<tbody>
					{#each users as user}
						<tr>
							<td>{user.email}</td>
							<td>{user.name}</td>
							<td>
								{#each user.roles as role }
									<span class="badge">{roleName(role)}</span>
								{/each}
							</td>
							<td class="actions-cell">
								<div class="action-icons">
									<a href="/system/users/{user.id}/edit" use:link aria-label="edit">
										<svg class="action-icon edit-icon" width="18" height="18">
											<use href="#fa-edit-icon" />
										</svg>
									</a>
								</div>
							</td>
						</tr>
					{/each }
				</tbody>
			</table>
		</div>
	{/snippet}
</Layout>
	
<style lang="scss">
	.content {
		background-color: #fff;
		min-height: 400px;
		box-shadow: 1px 0px 2px rgba(0, 0, 0, 0.1);
	}
</style>