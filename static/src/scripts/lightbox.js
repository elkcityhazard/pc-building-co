import PhotoSwipeLightbox from 'photoswipe/lightbox';
// css is included via postcss build process now

const lightbox = new PhotoSwipeLightbox({
    // may select multiple "galleries"
    gallery: '#galleryWork',
  
    // Elements within gallery (slides)
    children: 'a',
  
    // setup PhotoSwipe Core dynamic import
    pswpModule: () => import('photoswipe')
});
  

export const InitLightBox = function () {
    lightbox.init()
}
