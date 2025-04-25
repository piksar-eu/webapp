<script>
	import Breadcrumbs from "../../components/Breadcrumbs.svelte";
	import Layout from "../../components/Layout.svelte";
	import { getRoles, deleteRole as deleteRoleApiCall } from "../../api";
    import { link } from "svelte5-router"; 
    import { onMount } from "svelte";
	import { alert } from '../../shared';

	let roles = $state()
	onMount(async () => {
		roles = await getRoles()
	})

	const deleteRole = async (id) => {
		await deleteRoleApiCall(id)
		const rName = roles.find(r => r.id == id).name
		alert.set({type: "success", msg: `Usunęto rolę <strong>${rName}</strong>`});

		roles = roles.filter(r => r.id !== id)
	}
</script>

<Layout>
	{#snippet header()}
		<Breadcrumbs items={[{href: '/system', name: 'System'}, {name: 'Role'}]}></Breadcrumbs>
		<div class="header-buttons">
			<a href="/system/role/add" use:link class="button">Dodaj</a>
		</div>
	{/snippet}
	{#snippet main()}
	<div class="content grid">
		<svg xmlns="http://www.w3.org/2000/svg" style="display: none;">
			<!-- Font Awesome Edit Icon -->
			<symbol id="fa-edit-icon" viewBox="0 0 512 512">
			<path d="M441 58.9L453.1 71c9.4 9.4 9.4 24.6 0 33.9L424 134.1 377.9 88 407 58.9c9.4-9.4 24.6-9.4 33.9 0zM209.8 256.2L344 121.9 390.1 168 255.8 302.2c-2.9 2.9-6.5 5-10.4 6.1l-58.5 16.7 16.7-58.5c1.1-3.9 3.2-7.5 6.1-10.4zM373.1 25L175.8 222.2c-8.7 8.7-15 19.4-18.3 31.1l-28.6 100c-2.4 8.4-.1 17.4 6.1 23.6s15.2 8.5 23.6 6.1l100-28.6c11.8-3.4 22.5-9.7 31.1-18.3L487 138.9c28.1-28.1 28.1-73.7 0-101.8L474.9 25C446.8-3.1 401.2-3.1 373.1 25zM88 64C39.4 64 0 103.4 0 152L0 424c0 48.6 39.4 88 88 88l272 0c48.6 0 88-39.4 88-88l0-112c0-13.3-10.7-24-24-24s-24 10.7-24 24l0 112c0 22.1-17.9 40-40 40L88 464c-22.1 0-40-17.9-40-40l0-272c0-22.1 17.9-40 40-40l112 0c13.3 0 24-10.7 24-24s-10.7-24-24-24L88 64z"/>
			</symbol>
			
			<!-- Font Awesome Delete Icon -->
			<symbol id="fa-delete-icon" viewBox="0 0 448 512">
			<path d="M135.2 17.7C140.6 6.8 151.7 0 163.8 0L284.2 0c12.1 0 23.2 6.8 28.6 17.7L320 32l96 0c17.7 0 32 14.3 32 32s-14.3 32-32 32L32 96C14.3 96 0 81.7 0 64S14.3 32 32 32l96 0 7.2-14.3zM32 128l384 0 0 320c0 35.3-28.7 64-64 64L96 512c-35.3 0-64-28.7-64-64l0-320zm96 64c-8.8 0-16 7.2-16 16l0 224c0 8.8 7.2 16 16 16s16-7.2 16-16l0-224c0-8.8-7.2-16-16-16zm96 0c-8.8 0-16 7.2-16 16l0 224c0 8.8 7.2 16 16 16s16-7.2 16-16l0-224c0-8.8-7.2-16-16-16zm96 0c-8.8 0-16 7.2-16 16l0 224c0 8.8 7.2 16 16 16s16-7.2 16-16l0-224c0-8.8-7.2-16-16-16z"/>
			</symbol>
		</svg>
		<table>
			<thead>
			<tr>
				<th scope="col">Nazwa</th>
				<th scope="col">Uprawnienia</th>
				<th scope="col"></th>
			</tr>
			</thead>
			<tbody>
				{#each roles as role}
					<tr>
						<td>{role.name}</td>
						<td>
							{#each role.permissions as permission }
								<span class="badge">{permission}</span>
							{/each}
						</td>
						<td class="actions-cell">
							<div class="action-icons">
								<a href="/system/role/{role.id}/edit" use:link aria-label="edit">
									<svg class="action-icon edit-icon" width="18" height="18">
										<use href="#fa-edit-icon" />
									</svg>
								</a>
								<button onclick={() => {deleteRole(role.id) }} aria-label="delete">
									<svg class="action-icon delete-icon" width="18" height="18">
									<use href="#fa-delete-icon" />
									</svg>
								</button>
							</div>
							</td>
					</tr>
				{/each}
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