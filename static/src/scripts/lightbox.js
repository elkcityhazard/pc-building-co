import PhotoSwipeLightbox from "../../../node_modules/photoswipe/dist/photoswipe-lightbox.esm"
import PhotoSwipe from "../../../node_modules/photoswipe/dist/photoswipe.esm.js"
import "../../../node_modules/photoswipe/dist/photoswipe.css"
const lightbox = new PhotoSwipeLightbox({
    // may select multiple "galleries"
    gallery: '#galleryWork',
  
    // Elements within gallery (slides)
    children: 'a',
  
    // setup PhotoSwipe Core dynamic import
    pswpModule: PhotoSwipe
  });
  

export const InitLightBox = function() {
    lightbox.init()
}
