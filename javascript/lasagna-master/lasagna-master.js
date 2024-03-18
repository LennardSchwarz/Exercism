/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */
export function cookingStatus(time) {
    if (time === 0) {
        return 'Lasagna is done.';
    } else if (time > 0) {
        return 'Not done, please wait.';
    } else {
        return 'You forgot to set the timer.';
    }
}

export function preparationTime(layers, avgPrepMinsPerLayer = 2) {
    return layers.length * avgPrepMinsPerLayer
}

export function quantities(layers) {
    const CountNoodleLayers = layers.filter(x => x === "noodles").length;
    const CountSauceLayers = layers.filter(x => x === "sauce").length; 
    return {noodles: CountNoodleLayers * 50, sauce: CountSauceLayers* 0.2 }
}

export function addSecretIngredient(friendsList, myList) {
    myList.push(friendsList[friendsList.length - 1])
}

export function scaleRecipe(recipe, numPortions) {
    const scaledRecipeObj = {}
    const multiplier = numPortions/2
    for (let ingredient in recipe) {
        scaledRecipeObj[ingredient] = recipe[ingredient] * multiplier;
    }
    return scaledRecipeObj
}