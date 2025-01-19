# number_guessing_game.rb

# Generate a random number between 1 and 100
random_number = rand(1..100)

puts "Welcome to the Number Guessing Game!"
puts "I've picked a random number between 1 and 100."
puts "Can you guess what it is? You have unlimited attempts!"

loop do
  print "Enter your guess: "
  guess = gets.to_i

  if guess < random_number
    puts "Too low! Try again."
  elsif guess > random_number
    puts "Too high! Try again."
  else
    puts "Congratulations! You guessed the number #{random_number}!"
    break
  end
end

puts "Thanks for playing!"
