<script>
	import Breadcrumbs from "../../components/Breadcrumbs.svelte";
	import Layout from "../../components/Layout.svelte";
	import { getRole, getPermissions, saveRole as saveRoleApiCall } from "../../api";
    import { onMount } from "svelte";
    import { link, navigate } from "svelte-routing";
	import { alert } from '../../shared';

	let { id } = $props();

	let role = $state()
	let roleName = $state()
	let rolePermissions = $state()

	onMount(async () => {
		if (id) {
			role = await getRole(id)
			roleName = role.name
			rolePermissions = role.permissions
		}
	})

	const saveRole = async (e) => {
		e.preventDefault()

		try {
			await saveRoleApiCall(id, roleName, rolePermissions)
			
            navigate("/system/roles", { replace: true });
            alert.set({type: "success", msg: `Zapisano rolę <strong>${roleName}</strong>`});
		} catch (err) {
            alert.set({type: "error", msg: "Błąd podczas zapisu"});
		}
	}
</script>

<Layout>
	{#snippet header()}
		<Breadcrumbs items={[{href: '/system', name: 'System'}, {href: '/system/roles', name: 'Role'}, {name: role?.name ?? "Nowa rola"}]}></Breadcrumbs>
	{/snippet}
	{#snippet main()}
	<div class="content form">
		<form onsubmit={saveRole}>
			<div class="form-group">
				<label for="name" class="form-label">Nazwa</label>
				<input type="text" id="name" class="form-control" placeholder="Nazwa roli" bind:value={roleName}>
			</div>
			
			<div class="form-group">
				<div class="form-label">Uprawnienia</div>
				<div class="permissions-list">
				{#await getPermissions() then permissions }
					{#each permissions as permission }
						<div class="permission-item">
							<input type="checkbox" name="permissions[]" value="{permission.id}" id="permission-{permission.id}" bind:group={rolePermissions} class="permission-checkbox">
							<label for="permission-{permission.id}" class="permission-label">
								{permission.id}
								<div class="permission-description">{permission.name}</div>
							</label>
						</div>
					{/each}
				{/await}
				</div>
			</div>
			
			<div class="form-actions">
				<a use:link href="/system/roles" class="button button-secondary">Anuluj</a>
				<button type="submit" class="button button-primary">Zapisz</button>
			</div>
		</form>
	</div>
	{/snippet}
</Layout>
	
<style lang="scss">
	.content {
		background-color: #fff;
		min-height: 400px;
		box-shadow: 1px 0px 2px rgba(0, 0, 0, 0.1);
		padding: 20px;
	}


    .permissions-list {
      border: 1px solid #e9ecef;
      border-radius: 4px;
      overflow: hidden;
    }

    .permission-item {
      padding: 1rem;
      border-bottom: 1px solid #e9ecef;
      display: flex;
      align-items: center;
    }

    .permission-item:last-child {
      border-bottom: none;
    }

    .permission-item:nth-child(even) {
      background-color: #f8f9fa;
    }

    .permission-item:hover {
      background-color: #f1f3f5;
    }

    .permission-checkbox {
      margin-right: 1rem;
      width: 18px;
      height: 18px;
      cursor: pointer;
    }

    .permission-label {
      flex-grow: 1;
      cursor: pointer;
    }
    .permission-description {
      font-size: 0.8rem;
      color: #6c757d;
      margin-top: 0.25rem;
    }

</style>