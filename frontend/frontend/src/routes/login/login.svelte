<script>
  let username = '';
  let password = '';

  function logIn() {
    console.log(username, password);
  }


</script>

<div class="page">
  <div class="login">
    <h1>Login</h1>
    <input placeholder="Username" bind:value={username}/>
    <input placeholder="Password" type="password" bind:value={password} />

    <button on:click={logIn}>
        Sign In
    </button>
  </div>
</div>

<style>
  .page {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .login {
    width: 300px;
    padding: 20px;
    border: 1px solid rgb(255, 255, 255);
  }
</style>