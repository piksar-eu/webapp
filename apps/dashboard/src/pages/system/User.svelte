<script>
	import Breadcrumbs from "../../components/Breadcrumbs.svelte";
	import Layout from "../../components/Layout.svelte";
	import { getUser, getRoles, saveUser as saveUserApiCall } from "../../api";
    import { onMount } from "svelte";
    import { link, navigate } from "svelte-routing";
    import { alert } from "../../shared";

	let { id } = $props();

	let user = $state()
	let userName = $state()
	let userRoles = $state()
	let roles = $state()

	onMount(async () => {
		roles = await getRoles()
		if (id) {
			user = await getUser(id)

			userName = user.name
			userRoles = user.roles
		}
	})

	const saveUser = async (e) => {
		e.preventDefault()

		try {
			await saveUserApiCall(id, userName, userRoles)
			
            navigate("/system/users", { replace: true });
            alert.set({type: "success", msg: `Zapisano użytkownika <strong>${userName}</strong>`});
		} catch (err) {
            alert.set({type: "error", msg: "Błąd podczas zapisu"});
		}
	}
</script>

<Layout>
	{#snippet header()}
		<Breadcrumbs items={[{href: '/system', name: 'System'}, {href: '/system/users', name: 'Użytkownicy'}, {name: (user?.name && user.name != "" ) ? user.name : user?.email}]}></Breadcrumbs>
	{/snippet}
	{#snippet main()}
	<div class="content form">
		<form onsubmit={saveUser}>
			<div class="form-group">
				<label for="name" class="form-label">Email</label>
				<input type="text" id="name" class="form-control" placeholder="User email" value={user?.email} readonly disabled>
			</div>
			
			<div class="form-group">
				<label for="name" class="form-label">Nazwa</label>
				<input type="text" id="name" class="form-control" placeholder="User name" bind:value={userName}>
			</div>
			
			<div class="form-group">
				<div class="form-label">Role</div>
				<div class="permissions-list">
					{#each roles as role }
						<div class="role-item">
							<input type="checkbox" name="roles[]" value="{role.id}" id="role-{role.id}" bind:group={userRoles} class="role-checkbox">
							<label for="role-{role.id}" class="role-label">
								{role.name}
								<div class="role-description">
									{#each role.permissions as permission }
										<span class="badge">{permission}</span>
									{/each}
								</div>
							</label>
						</div>
					{/each}
				</div>
			</div>
			
			<div class="form-actions">
				<a use:link href="/system/users" class="button button-secondary">Anuluj</a>
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


    .roles-list {
      border: 1px solid #e9ecef;
      border-radius: 4px;
      overflow: hidden;
    }

    .role-item {
      padding: 1rem;
      border-bottom: 1px solid #e9ecef;
      display: flex;
      align-items: center;
    }

    .role-item:last-child {
      border-bottom: none;
    }

    .role-item:nth-child(even) {
      background-color: #f8f9fa;
    }

    .role-item:hover {
      background-color: #f1f3f5;
    }

    .role-checkbox {
      margin-right: 1rem;
      width: 18px;
      height: 18px;
      cursor: pointer;
    }

    .role-label {
      flex-grow: 1;
      cursor: pointer;
    }
    .role-description {
      font-size: 0.8rem;
      color: #6c757d;
      margin-top: 0.25rem;
    }

</style>