<script lang="ts">
  import { Greet } from "../wailsjs/go/main/App.js";

  let resultText: string = "Please enter your name below 👇";
  let name: string;

  function greet(): void {
    Greet(name).then((result) => (resultText = result));
  }
</script>

<main>
  <div class="center">
    <img alt="Wails logo" id="logo" src="./logo-universal.png" />
    <div class="result" id="result">{resultText}</div>
    <div class="input-box" id="input">
      <input
        autocomplete="off"
        bind:value={name}
        class="input"
        id="name"
        type="text"
      />
      <button class="btn" on:click={greet}>Greet</button>
    </div>
  </div>
</main>

<style>
  main {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .center {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
  }

  #logo {
    display: block;
    width: 200px;
    height: auto;
    margin: 0 auto 1.5rem;
    object-fit: contain;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 0 auto 1.5rem;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
