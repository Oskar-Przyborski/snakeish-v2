<div class="container">
	<slot />
</div>

<style lang="scss">
	.container {
        margin: auto;
        height:100%;
        width:100%;
	}
    @media (min-width: 576px) {
        .container{
            max-width: 540px;
        }
    }
    @media (min-width: 768px) {
        .container{
            max-width: 720px;
        }
    }
    @media (min-width: 992px) {
        .container{
            max-width: 960px;
        }
    }
	@media (min-width: 1200px) {
        .container{
            max-width: 1110px;
        }
	}
</style>
