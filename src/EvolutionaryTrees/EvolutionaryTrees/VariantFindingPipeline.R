# This R script takes a distance matrix between genomes as input.
# It produces a PCoA plot of the distances between genomes.

# Import needed libraries. Please install these in R beforehand using install.packages("package_name")

library(ggcorrplot)
library(reshape)
library(stringr)
library(ggplot2)
library(ape) #ape library to compute PCoA of our matrix

# Now set working directory. This should be wherever the files are stored and is the only line that the user needs to edit.

# Use "Session" --> "Set Working Directory" --> "To Source File Location"

#PLOT 1: Generating a PCoA plot of the data between 2020 and 2024

# The file was formed by first sampling 20 genomes every two weeks from the UK, and then taking the Jaccard distance between every pair of k-mer frequency maps for these genomes.

table <- read.csv(file="Matrices/JaccardBetaDiversityMatrix_2020-2024.csv")

table <- table[-c(1)] # trim out first column

table <- table[, -ncol(table)] # get rid of weird final column

matrix <- as.matrix(table)

pcoa_data <- pcoa(matrix, correction="none", rn=NULL)

pcoa_vectors <- data.frame(pcoa_data$vectors)

ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2)) + geom_point()

ggsave("Plots/JaccardPCoA_2020-2024.png")

#PLOT 2: Generating a PCoA plot of the data between 2020 and 2022
# Read in the file and process the table.
table <- read.csv(file="Matrices/JaccardBetaDiversityMatrix_2020-2022.csv")

#trim the first column out because it only contains names
table <- table[-c(1)]
table <- table[, -c(3454)] # trim out weird extra column at end of the matrix file

matrix <- as.matrix(table)

pcoa_data <- pcoa(matrix, correction="none", rn=NULL) #This step may take a minute or two
pcoa_vectors <- data.frame(pcoa_data$vectors)
# columns contains a vector for each point after PCoA tries to assign data points to vectors to preserve distances between points.

colnames(table)


# Now, plot the data
ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2)) + geom_point()
ggsave("Plots/JaccardPCoA_2020-2022.png")


#PLOT 3: Using the matrix labelled using our variant classifier to make a colored PCoA plot for 2020-2024
# Load and preprocess the data
table3 <- read.csv(file="Matrices/JaccardBetaDiversityMatrixLabelled_2020-2024.csv")

#trim the first column out because it only contains names
table3 <- table3[-c(1)]

# we want our table to be square
table3 <- table3[, -ncol(table3)] # get rid of weird final column

matrix3 <- as.matrix(table3)

pcoa_data <- pcoa(matrix3, correction="none", rn=NULL)
pcoa_vectors <- data.frame(pcoa_data$vectors)
# columns contains a vector for each point after PCoA tries to assign data points to vectors to preserve distances between points.

#split the colnames by "_" then keep the last entry, which is the variant name
Variant <- sapply(strsplit(colnames(table3), "_"), function(x) tail(x, 1))

cbind(pcoa_vectors, Variant) # adding column

# Now, plot the data, colored by variant
ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2, color=Variant)) + geom_point()

ggsave("Plots/JaccardPCoALabelled_2020-2024.png")


#PLOT 4: Using the matrix labelled using our variant classifier to make a colored PCoA plot
table4 <- read.csv(file="Matrices/JaccardBetaDiversityMatrixLabelled_2020-2022.csv")

#trim the first column out because it only contains names
table4 <- table4[-c(1)]

table4 <- table4[, -ncol(table4)] # get rid of weird final column

matrix4 <- as.matrix(table4)

pcoa_data <- pcoa(matrix4, correction="none", rn=NULL)
pcoa_vectors <- data.frame(pcoa_data$vectors)
# columns contains a vector for each point after PCoA tries to assign data points to vectors to preserve distances between points.

colnames(table4)


Variant <- sub(".*_.*_.*_.*_", "", colnames(table4))

cbind(pcoa_vectors, Variant) # adding column

# Now, plot the data, colored by variant
ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2, color=Variant)) + geom_point()
ggsave("Plots/JaccardPCoALabelled_2020-2022.png")

