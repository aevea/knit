<script>
  export let name;
  import OldestPR from "./components/OldestPR.svelte";
</script>

<style>
  :global(body) {
    background: lightblue;
  }

  main {
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
    text-align: center;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
  <h1>Hello {name}!</h1>
  <OldestPR />
</main>
