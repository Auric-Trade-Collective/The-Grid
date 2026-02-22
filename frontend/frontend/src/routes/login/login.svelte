<script>
  import gridDesign from "../../../../../mockups/gridlogin.png";

  let username = '';
  let password = '';

  function logIn() {
    console.log(username, password);
  }
</script>

<div class="page">
  <!-- LEFT SIDE -->
  <div class="left">
    <div class="login-card">
      <h1>LOGIN / SIGN UP</h1>

      <div class="input-wrapper">
        <input
          type="text"
          placeholder="Username"
          bind:value={username}
        />
      </div>

      <div class="input-wrapper">
        <input
          type="password"
          placeholder="Password"
          bind:value={password}
        />
      </div>

      <button on:click={logIn}>
        Sign In
      </button>
    </div>
  </div>

  <!-- RIGHT SIDE -->
  <div class="right">
    <img src={gridDesign} alt="Mockup visual" />
  </div>
</div>

<style>
  .page {
    height: 100vh;
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .left {
    background-color: #244e54;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .right {
    position: relative;
    overflow: hidden;
    width: 100%;
    object-fit: fill;
  }

  .right img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .login-card {
    width: 360px;
    padding: 3rem;
    border-radius: 24px;
    background-color: #c5ecf0;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  h1 {
    text-align: center;
    margin-bottom: 1rem;
  }

  .input-wrapper input {
    width: 100%;
    padding: 0.9rem 1.2rem;
    border-radius: 999px;
    border: none;
    outline: none;
  }

  button {
    margin-top: 1rem;
    padding: 1rem;
    border-radius: 999px;
    border: none;
    cursor: pointer;
    background-color: rgb(44, 201, 226);
  }
</style>