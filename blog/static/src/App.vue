<template>
  <div id="app">
    <!-- Start of nav bar 
		============================================= -->
		<nav id="poppin-nav">
			<div id="nav-off"></div>
			<div id="header-logo-1" class="text-center">
				<a href="#"><img src="./assets/images/logo.png" alt="img"></a>
			</div>
			<ul id="menu">
				<li class="has-submenu text-capitalize">
					<a href="index.html">home</a>
				</li>
				<li class="has-submenu">
					<a href="#">Blog <span class="ti-angle-down"></span></a>
					<ul class="submenu">
						<li><a href="blog.html">Blog Grid</a></li>
						<li><a href="blog-single.html">Single Post</a></li>
					</ul>
				</li>
				<li><a href="contact.html">Contact</a></li>
			</ul>
		</nav>
		<div id="menu-overlay"></div>
	<!-- End of nav bar 
		============================================= -->



	<!-- Start of header
		============================================= -->
		<header id="site-header" class="not-stuck">
			<div class="container">
				<div class="row">
					<div id="header-logo">
						<a href="#"><img src="./assets/images/logo.png" alt="img"></a>
					</div>

					<div id="menu-burger" class="pull-right not-stuck">
						<div id="menu-bar">
							<span class="icon-bar top"></span>
							<span class="icon-bar middle"></span>
							<span class="icon-bar bottom"></span>
						</div>
					</div>
				</div>
			</div>
		</header>
	<!-- End of header
		============================================= -->





	<!-- Start of page head section
		============================================= -->
		<section id="page-head" class="page-head-section">
			<div class="page-head-overlay"></div>
			<div id="page-head-effect" class="page-head-effect">
				<canvas id="demo-canvas"></canvas>
			</div>
			<!-- // page-head-style -->
			<div class="container">
				<div class="row">
					<div class="banner-content">
						<div class="page-head-content">
							<div class="page-head-title text-capitalize pb40">
								<h1>This is Puppin</h1>
								<span class="descrip">Creative Designer</span>
							</div>
							<!-- //title -->

							<div class="page-head-social">
								<div  class="social-title text-uppercase">
									<span>socials</span>
								</div>
								<!-- //title -->
								<div class="page-head-social-item ul-li">
									<ul class="page-head-social-list">
										<li><a href="#"><span class="ti-facebook"></span></a></li>
										<li><a href="#"><span class="ti-twitter-alt"></span></a></li>
										<li><a href="#"><span class="ti-google"></span></a></li>
										<li><a href="#"><span class="ti-instagram"></span></a></li>
									</ul><!-- /.page-head-social-list -->
								</div><!-- /.page-head-social-item -->
							</div>
							<!-- /.page-head-social -->
						</div>
						<!-- /.page-head-content -->
					</div>
				</div><!--  /.row-->
			</div><!--  /.container -->
		</section>
	<!-- End of page head section
		============================================= -->





	<div class="tlinks">Collect from <a href="http://www.cssmoban.com/" >自助建站</a></div>


	<!-- Start of newsletter section
		============================================= -->
		<section id="newsletter" class="newsletter-section">
			<div class="page-head-overlay"></div>
			<div class="container">
				<div class="row">
					<div class="newsletter-section-content">
						<div class="newsletter-title pb75">
							<h2 class="deep-black">Subscribe On My Newsletter</h2>
						</div>

						<div class="newsletter-form">
							<form action="#" class="input-group">
								<input type="email" class="form-control" placeholder="Enter Your Email">
								<span class="input-group-addon">
									<button type="submit"><span class="deep-black ti-arrow-right"></span></button>  
								</span>
							</form>
						</div>
					</div>
					<!-- //.newsletter-section-content -->

				</div><!--  /.row -->
			</div><!--  /.conatainer -->
		</section>
	<!-- End of newsletter section
		============================================= -->

	<!-- Start of footer section
		============================================= -->
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
	<!-- End of footer section
		============================================= -->
  </div>
</template>

<script>

import 'bootstrap/dist/css/bootstrap.min.css';
import '@icon/themify-icons/themify-icons.css';
import 'owl.carousel/dist/assets/owl.carousel.css';

// import Nav from './components/Nav.vue'
// import Footer from './components/Footer.vue'

export default {
  name: 'app',
  components: {
    
  }
}
</script>

<style>
@import url('./assets/css/menu.css');
@import url('./assets/css/style.css');
</style>
