let { clientWidth = 0 } = document.documentElement

export const Parallax = function () {
  const bgImage = document.getElementById("homeHeader");
  if (!bgImage || clientWidth < 968) {
    return;
  }


  const updateParallax = () => {
    const currentScroll = document.documentElement.scrollTop;
    clientWidth = document.documentElement.clientWidth

    let p = 50 - currentScroll / 10;
    bgImage.style.backgroundPositionY = p + "%";
    updateSize(bgImage)
  };

  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        window.addEventListener("scroll", throttle(updateParallax, 16.67));
        updateParallax(); // Initial call to set the position
      } else {
        window.removeEventListener("scroll", throttle(updateParallax, 16.67));
      }
    });
  });
  updateParallax()
  observer.observe(bgImage);
};

const calculateScroll = () => {
  const documentHeight = document?.documentElement?.scrollHeight
  const currentScroll = document?.documentElement.scrollTop
  const result = currentScroll / documentHeight
  return result
}

const updateSize = function (el) {
  const perc = calculateScroll()
  el.style.backgroundSize = (1.20 + perc) * 100 + '%'

}

export function throttle(cb = null, delay = 100) {
  if (!cb) return

  let shouldWait = false
  let waitingArgs

  const timeoutFunc = () => {
    if (waitingArgs == null) {
      shouldWait = false
    } else {
      cb(...waitingArgs)
      waitingArgs = null
      setTimeout(timeoutFunc, delay)
    }
  }

  return (...args) => {
    if (shouldWait) return
    cb(...args)
    shouldWait = true 
    setTimeout(timeoutFunc, delay)
  }


}