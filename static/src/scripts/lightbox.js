import PhotoSwipeLightbox from 'photoswipe/lightbox';
// css is included via postcss build process now

const lightbox = new PhotoSwipeLightbox({
    gallery: '#galleryWork',
    children: 'a',
    pswpModule: () => import('photoswipe')
});
  

export const InitLightBox = function () {
    lightbox.init()
}
