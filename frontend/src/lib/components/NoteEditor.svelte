<script lang="ts">
 export let summary="";
 export let keywords:string[]=[];
</script>

<div>
 <p>{summary}</p>
 <ul>
  {#each keywords as k}
   <li>{k}</li>
  {/each}
 </ul>
</div>