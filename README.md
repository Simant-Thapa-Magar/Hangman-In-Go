<h1>Hangman using Go</h1>

<h3>Game Algorithm</h3>

<p>
<ol class="ak-ol" data-indent-level="1"><li><p data-renderer-start-pos="1090">Initialize game with defining variable to store game’s state as 0 in the form of <code class="code css-1xivnph" data-renderer-mark="true">hangmanState</code></p></li><li><p data-renderer-start-pos="1187">Ask input from user for game difficulty until valid selection is made</p><ol class="ak-ol" data-indent-level="2"><li><p data-renderer-start-pos="1260">1 for easy difficulty words</p></li><li><p data-renderer-start-pos="1291">2 for medium difficulty words</p></li><li><p data-renderer-start-pos="1324">3 for hard difficulty words</p></li></ol></li><li><p data-renderer-start-pos="1357">Based on user selection randomly pick a word for the game to proceed</p></li><li><p data-renderer-start-pos="1429">Declare necessary variables to track different aspects of game such as</p><ol class="ak-ol" data-indent-level="2"><li><p data-renderer-start-pos="1503"><code class="code css-1xivnph" data-renderer-mark="true">wordGuessTrack</code> to keep track of words guessed</p></li><li><p data-renderer-start-pos="1552"><code class="code css-1xivnph" data-renderer-mark="true">gameEnd</code> to acknowledge if game is in progress or finished</p></li><li><p data-renderer-start-pos="1613"><code class="code css-1xivnph" data-renderer-mark="true">showHint</code> to acknowledge if user has asked to see the hint and display it</p></li><li><p data-renderer-start-pos="1689"><code class="code css-1xivnph" data-renderer-mark="true">success</code> to acknowledge if game ended with user as winner or not</p></li><li><p data-renderer-start-pos="1756"><code class="code css-1xivnph" data-renderer-mark="true">oldMessage</code> to keep track of any temporary message to be shown</p></li></ol></li><li><p data-renderer-start-pos="1823">Display game’s state in terms of hangman state, word to be guessed, letters displayed and any messaged</p><ol class="ak-ol" data-indent-level="2"><li><p data-renderer-start-pos="1929">Clear everything on screen</p></li><li><p data-renderer-start-pos="1959">Based on <code class="code css-1xivnph" data-renderer-mark="true">hangmanState</code> retrieve content from specific file on states folder and display</p></li><li><p data-renderer-start-pos="2049">Display correct letter for those already guessed or displayed initially and display _ for letters to be guessed</p></li><li><p data-renderer-start-pos="2164">If <code class="code css-1xivnph" data-renderer-mark="true">showHint</code> is set to true the display hint associated with word</p></li><li><p data-renderer-start-pos="2232">If there’s any <code class="code css-1xivnph" data-renderer-mark="true">oldMessaged</code> then display it</p></li></ol></li><li><p data-renderer-start-pos="2280">If <code class="code css-1xivnph" data-renderer-mark="true">gameEnd</code> is set to true go to step 12</p></li><li><p data-renderer-start-pos="2323">Ask User input</p></li><li><p data-renderer-start-pos="2341">If input is “hint” or has a length of 1 then procced to step 9 else go to step 7</p></li><li><p data-renderer-start-pos="2425">Read user input</p><ol class="ak-ol" data-indent-level="2"><li><p data-renderer-start-pos="2444">If user has typed <code class="code css-1xivnph" data-renderer-mark="true">hint</code> then set <code class="code css-1xivnph" data-renderer-mark="true">showHint</code> to true</p></li><li><p data-renderer-start-pos="2496">Else determine if what type of guess user has made</p><ol class="ak-ol" data-indent-level="3"><li><p data-renderer-start-pos="2550">If the guessed letter is already present in <code class="code css-1xivnph" data-renderer-mark="true">wordGuessTrack</code> then its already guessed previously</p><ol class="ak-ol" data-indent-level="4"><li><p data-renderer-start-pos="2648">Set message info to display letter was already guessed before</p></li><li><p data-renderer-start-pos="2713">Return true with message</p></li></ol></li><li><p data-renderer-start-pos="2743">Else check if word to be guessed contains the letter just guessed</p><ol class="ak-ol" data-indent-level="4"><li><p data-renderer-start-pos="2812">If it contains then add the letter to <code class="code css-1xivnph" data-renderer-mark="true">wordGuessTrack</code> </p></li><li><p data-renderer-start-pos="2869">Return true</p></li></ol></li><li><p data-renderer-start-pos="2886">If its incorrect guess the return false</p></li></ol></li><li><p data-renderer-start-pos="2931">If the guess was incorrect then update the <code class="code css-1xivnph" data-renderer-mark="true">hangmanState</code> by incrementing by 1</p></li><li><p data-renderer-start-pos="3011">If any message was received then set it to <code class="code css-1xivnph" data-renderer-mark="true">oldMessage</code></p></li></ol></li><li><p data-renderer-start-pos="3070">Check if <code class="code css-1xivnph" data-renderer-mark="true">hangmanState</code> is in final state</p><ol class="ak-ol" data-indent-level="2"><li><p data-renderer-start-pos="3113">If <code class="code css-1xivnph" data-renderer-mark="true">hangmanState</code> is equal to 9 then set <code class="code css-1xivnph" data-renderer-mark="true">gameOver</code> to true</p></li><li><p data-renderer-start-pos="3172">Else check if all letters have been guessed correctly</p><ol class="ak-ol" data-indent-level="3"><li><p data-renderer-start-pos="3229">If true then set <code class="code css-1xivnph" data-renderer-mark="true">gameOver</code> to true and <code class="code css-1xivnph" data-renderer-mark="true">success</code> to true</p></li></ol></li></ol></li><li><p data-renderer-start-pos="3290">Go to step 5</p></li><li><p data-renderer-start-pos="3306">Display game over info</p><ol class="ak-ol" data-indent-level="2"><li><p data-renderer-start-pos="3332">If <code class="code css-1xivnph" data-renderer-mark="true">success</code> user won</p></li><li><p data-renderer-start-pos="3355">Else user lost</p></li></ol></li></ol>
</p>