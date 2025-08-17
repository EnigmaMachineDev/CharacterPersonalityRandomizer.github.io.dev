
        const generateButton = document.getElementById('generateButton');
        const characterDisplay = document.getElementById('characterDisplay');

        let character = {};
        let namesData = null;
        let personalityData = null;

        const getRandomElement = (arr) => arr[Math.floor(Math.random() * arr.length)];

        const generateIdentity = () => {
            const sex = getRandomElement(['Male', 'Female']);
            const firstName = sex === 'Male'
                ? getRandomElement(namesData.firstName.male)
                : getRandomElement(namesData.firstName.female);
            const lastName = getRandomElement(namesData.lastNames);
            return { firstName, lastName, sex };
        };

        const generatePersonality = () => {
            const alignment = getRandomElement(personalityData.alignment);
            const personalityType = getRandomElement(personalityData.personalityType);
            return { alignment, personalityType };
        };

        const renderCharacter = () => {
            if (Object.keys(character).length === 0) {
                characterDisplay.innerHTML = '<p>Click the button to generate a character!</p>';
                return;
            }

            const isMobile = window.innerWidth <= 768;
            
            if (isMobile) {
                characterDisplay.innerHTML = `
                    <div class="section-with-reroll">
                        <div class="section-header">
                            <strong class="section-header-text">Identity</strong>
                        </div>
                        <div class="details-grid">
                            <p><strong>Name:</strong></p><p>${character.firstName} ${character.lastName}</p>
                            <p><strong>Sex:</strong></p><p>${character.sex}</p>
                        </div>
                        <button class="reroll-button" id="rerollIdentityBtn">↻</button>
                    </div>
                    <div class="section-divider"></div>
                    <div class="section-with-reroll">
                        <div class="section-header">
                            <strong class="section-header-text">Personality</strong>
                        </div>
                        <div class="details-grid">
                            <p><strong>Alignment:</strong></p><p>${character.alignment}</p>
                            <p><strong>Personality Type:</strong></p><p>${character.personalityType.name} <a href="${character.personalityType.link}" target="_blank">(link)</a></p>
                        </div>
                        <button class="reroll-button" id="rerollPersonalityBtn">↻</button>
                    </div>
                `;
            } else {
                characterDisplay.innerHTML = `
                    <button class="reroll-button" id="rerollIdentityBtn">↻</button>
                    <div id="identitySection">
                        <div class="section-header">
                            <strong class="section-header-text">Identity</strong>
                        </div>
                        <div class="details-grid">
                            <p><strong>Name:</strong></p><p>${character.firstName} ${character.lastName}</p>
                            <p><strong>Sex:</strong></p><p>${character.sex}</p>
                        </div>
                    </div>
                    <div class="section-divider"></div>
                    <button class="reroll-button" id="rerollPersonalityBtn">↻</button>
                    <div id="personalitySection">
                        <div class="section-header">
                            <strong class="section-header-text">Personality</strong>
                        </div>
                        <div class="details-grid">
                            <p><strong>Alignment:</strong></p><p>${character.alignment}</p>
                            <p><strong>Personality Type:</strong></p><p>${character.personalityType.name} <a href="${character.personalityType.link}" target="_blank">(link)</a></p>
                        </div>
                    </div>
                `;
            }

            document.getElementById('rerollIdentityBtn').addEventListener('click', handleRerollIdentity);
            document.getElementById('rerollPersonalityBtn').addEventListener('click', handleRerollPersonality);
        };

        const handleGenerateCharacter = () => {
            if (!namesData || !personalityData) return;
            characterDisplay.innerHTML = '<p>Generating character...</p>';
            const identity = generateIdentity();
            const personality = generatePersonality();
            character = { ...identity, ...personality };
            renderCharacter();
        };

        const handleRerollIdentity = () => {
            if (!namesData) return;
            const identity = generateIdentity();
            character.firstName = identity.firstName;
            character.lastName = identity.lastName;
            character.sex = identity.sex;
            renderCharacter();
        };

        const handleRerollPersonality = () => {
            if (!personalityData) return;
            const personality = generatePersonality();
            character.alignment = personality.alignment;
            character.personalityType = personality.personalityType;
            renderCharacter();
        };

        window.addEventListener('DOMContentLoaded', async () => {
            generateButton.disabled = true;
            characterDisplay.innerHTML = '<p>Loading data...</p>';
            try {
                const [namesResponse, personalityResponse] = await Promise.all([
                    fetch('names.json'),
                    fetch('randomizer.json')
                ]);
                namesData = await namesResponse.json();
                personalityData = await personalityResponse.json();

                generateButton.disabled = false;
                characterDisplay.innerHTML = '<p>Click the button to generate a character!</p>';
            } catch (error) {
                characterDisplay.innerHTML = '<p style="color: red;">Failed to load character data. Please refresh the page.</p>';
                console.error('Error loading data:', error);
            }
        });

        generateButton.addEventListener('click', handleGenerateCharacter);
    