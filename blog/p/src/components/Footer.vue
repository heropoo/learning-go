<template>
  <footer>
    <div class="footer-area footer-2">
        <div class="container">
            <div class="row">
                <div class="copy-right-area text-center">
                    <div class="page-head-social-item ul-li">
                        <ul class="page-head-social-list">
                            <li><a href="#"><span class="ti-facebook"></span></a></li>
                            <li><a href="#"><span class="ti-twitter-alt"></span></a></li>
                            <li><a href="#"><span class="ti-google"></span></a></li>
                            <li><a href="#"><span class="ti-instagram"></span></a></li>
                        </ul><!-- /.page-head-social-list -->
                    </div>
                    <span>© puppin 2018 - All Rights Reserved - More Templates <a href="http://www.cssmoban.com/" target="_blank" title="模板之家">模板之家</a> - Collect from <a href="http://www.cssmoban.com/" title="网页模板" target="_blank">网页模板</a></span>
                </div>
                <!-- //copy-right-area -->
            </div><!--  /.container -->
        </div><!--  /.row-->
    </div><!--  /.footer-area -->
</footer>
</template>

<script>
export default {
  name: 'Footer',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.footer-area {
	padding: 50px 0px;
	max-width: 1170px;
	margin:  0 auto;
	position:  relative;
}
footer {
	position: relative;
}
.scrollup {
	position: absolute;
	bottom: 88px;
	right: 0%;
	display: none;
	transform: rotate(-90deg);
	border: 1px solid #090909;
	border-right: none;
	padding: 32px;
}
.scrollup.custom {
	bottom: 63px;
}
.scrollup:after {
	content: '';
	position: absolute;
	background-color: #090909;
	height: 1px;
	width: 19px;
	top: 10px;
	transform: rotate(-90deg);
	left: 54px;
}
.scrollup:before {
	content: '';
	position: absolute;
	background-color: #090909;
	height: 1px;
	width: 19px;
	transform: rotate(-90deg);
	right: -9px;
	bottom: 10px;
}
.back-top {
	position: absolute;
	width: 90px;
	top: 20px;
	letter-spacing: 1px;
}
.footer-2 {
	background-color: #f2f2f2;
	margin-top: 50px;
}
.footer-2 .page-head-social-item {
	position: inherit;
}
</style>