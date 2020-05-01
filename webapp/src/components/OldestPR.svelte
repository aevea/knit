<script>
  import Loader from "./Loader.svelte";
  import { PullRequestService } from "../services/client.js";
  import { PADDING } from "../styling/consts.js";

  const prService = new PullRequestService();

  async function getOldestPR() {
    let data = await prService.oldest("", {
      Repository: "fallion/fallion"
    });

    console.log(data);

    return data;
  }

  let PRData = getOldestPR();
</script>

<style>
  .container {
    background: white;
    max-width: 600px;
    max-height: 300px;
    border: 1px solid black;
    border-radius: 10px;
    padding: var(--padding);
    display: inline-block;
  }

  a[target="_blank"]::after {
    content: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAAQElEQVR42qXKwQkAIAxDUUdxtO6/RBQkQZvSi8I/pL4BoGw/XPkh4XigPmsUgh0626AjRsgxHTkUThsG2T/sIlzdTsp52kSS1wAAAABJRU5ErkJggg==);
    margin: 0 3px 0 5px;
  }
</style>

<div class="container" style="--padding: {PADDING}">
  <h2>Longest Open PR</h2>
  {#await PRData}
    <Loader />

  {:then data}
    <a href={data.uRL} target="_blank">{data.title}</a>

    <h3>Open for {data.openForDays} days</h3>

  {:catch error}
    <p>Something went wrong: {error.message}</p>
  {/await}

</div>
