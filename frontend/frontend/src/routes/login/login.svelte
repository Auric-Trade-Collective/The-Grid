<script>

  import pizza from "../../../../../mockups/pizza.jpg"
  let username = '';
  let password = '';

  function logIn() {
    console.log(username, password);
  }


</script>

<div class="page">
  <div class="centered">
    <div class="login">
      <h1>LOGIN/SIGN UP</h1>
      <input placeholder="Username" bind:value={username}/>
      <input placeholder="Password" type="password" bind:value={password} />

      <button on:click={logIn}>
          Sign In
      </button>
    </div>
  </div>
  <div class="centered">
    <img class="loginPizza" src={pizza} alt="Cant find image!" />
  </div>
</div>

<style>
  .page {
    height: 100vh;
    display: grid;
    grid-template-columns: auto auto auto;
  }

  .centered {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }

  .login {
    width: 300px;
    padding: 20px;
    border: 1px solid rgb(255, 255, 255);
  }

  .loginPizza {
    height: 100vh;
    width: auto;
  }

</style>