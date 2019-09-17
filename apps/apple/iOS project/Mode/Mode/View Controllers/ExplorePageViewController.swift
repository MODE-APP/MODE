//
//  ExplorePageViewController.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/16/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class ExplorePageViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        
    }
    

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */

}

//extension ProfileViewController: UICollectionViewDelegate, UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
//
//    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
////        return dataSource.count
//    }
//
//    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
////        guard let cell = profileCollectionView.dequeueReusableCell(withReuseIdentifier: "photoCell", for: indexPath) as? PostCollectionViewCell else {return UICollectionViewCell()}
////
////        cell.image = dataSource[indexPath.row]
////
////        return cell
//    }
//
//    // MARK: - FlowLayout
//
////    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
////        let size = CGSize(width: profileCollectionView.frame.width / 3 - 1, height: profileCollectionView.frame.width / 3 * 1.6 - 1)
////        return size
//    }
//
//    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, referenceSizeForFooterInSection section: Int) -> CGSize {
//        if loadedAllPosts {
//            return CGSize.zero
//        }
//        return CGSize(width: profileCollectionView.frame.width, height: 55)
//    }
//
//    func collectionView(_ collectionView: UICollectionView, viewForSupplementaryElementOfKind kind: String, at indexPath: IndexPath) -> UICollectionReusableView {
//        guard let footerView = profileCollectionView.dequeueReusableSupplementaryView(ofKind: kind, withReuseIdentifier: "loadDataFooter", for: indexPath) as? PostFooterActivtityIndicatorCollectionReusableView else {fatalError("Couldn't get footer view at \(#function)")}
//
//        return footerView
//    }
//
//    func scrollViewDidScroll(_ scrollView: UIScrollView) {
//        let heightFromBottom = scrollView.contentSize.height - scrollView.contentOffset.y
//        if heightFromBottom < 1300 && loadingPosts == false && loadedAllPosts == false{
//            loadMorePosts()
//            loadingPosts = true
//        }
//    }
//}
