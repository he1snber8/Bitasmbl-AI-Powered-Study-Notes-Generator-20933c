<script lang="ts">
 let content=""; let notes=null;
 async function generate(){ /* call backend */ }
</script>

<textarea bind:value={content}></textarea>
<button on:click={generate}>Generate</button>
{#if notes}
 <section></section>
{/if}