const myScript = {
    init: () => {
        myScript.mySal();
        myScript.myParty();
        myScript.myCarausel();
    },
    mySal: () => {
        sal();
    },
    myParty: () => {
        party.confetti(runButton, {
            count: party.variation.range(20, 40),
        });
    },
    myCarausel: () => {
        var myCarousel = document.querySelector('#carouselExampleDark')
        var carousel = new bootstrap.Carousel(myCarousel)
    }
};

myScript.init();