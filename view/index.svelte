<script>
  import Header from './Header.svelte'
  import Story from './Story.svelte'
  export let stories = []
</script>
<title>Hacker News Feed</title>
<Header />

{#each stories as story}
  <Story {story} />
{/each}

<!-- REMOVED TABLE BELOW -->

<!-- <table border="1" cellpadding="10">
  {#if stories.length > 0}
    <thead>
      {#each Object.keys(stories[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each stories as story}
    <tr>
      {#each Object.keys(story) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/${story.id || 0}`}>{story[key]}</a></td>
        {:else}
          <td>{story[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<br />

<a href={`/new`}>New Story</a>

<style>
  table {
    border-collapse: collapse;
  }
</style> --> -->
