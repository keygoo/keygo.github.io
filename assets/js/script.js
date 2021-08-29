const myScript = {
    init: () => {
        myScript.mySal();
        myScript.myParty();
    },
    mySal: () => {
        sal();
    },
    myParty: () => {
        party.confetti(runButton, {
            count: party.variation.range(20, 40),
        });
    }
};

myScript.init();